package oss_aliyun

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kysion/oss-library/internal/logic/osser"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
	"github.com/kysion/oss-library/oss_model/oss_do"
	"github.com/kysion/oss-library/oss_model/oss_entity"
	"github.com/kysion/oss-library/oss_model/oss_enum"
)

type sOssAliyun struct {
	osser.OSSer

	modules  oss_interface.IModules
	dao      *oss_dao.XDao
	category string
}

func NewOssAliyun(modules oss_interface.IModules) oss_interface.IOssAliyun {
	return &sOssAliyun{
		modules:  modules,
		dao:      modules.Dao(),
		category: "阿里云Oss存储",
	}
}

// CreateProvider 添加渠道商配置
func (s *sOssAliyun) CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error) {
	info.ProviderNo = oss_enum.Oss.Type.Aliyun.Code()
	info.ProviderName = oss_enum.Oss.Type.Aliyun.Description()
	// 或者还有其他拓展字段...

	provider, err := s.modules.OssServiceProviderConfig().CreateProvider(ctx, info)

	return provider, err
}

// RegisterBucket 添加Bucket存储空间配置
func (s *sOssAliyun) RegisterBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error) {
	info.ProviderNo = oss_enum.Oss.Type.Aliyun.Code() // aliyun

	result, err := s.modules.OssBucketConfig().CreateBucket(ctx, info)

	return result, err
}

// ---------------------------------------- 文件操作 ---------------------------------------------------

// UploadFile 上传文件
func (s *sOssAliyun) UploadFile(ctx context.Context, info *oss_model.PutObject) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectKey := info.ObjectKey
	// yourLocalFileName填写本地文件的完整路径。
	localFilePath := info.FilePath

	// 上传文件。
	err = bucket.PutObjectFromFile(objectKey, localFilePath)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_upload_file_failed}")
	}

	return true, nil
}

// PartPutUpload 断点续传文件
func (s *sOssAliyun) PartPutUpload(ctx context.Context, info *oss_model.PartPutObject) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// UploadFile(文件key，源文件，分片大小，并发数，是否开启断点续传，断点续传文件)
	err = bucket.UploadFile(info.ObjectKey, info.FilePath, info.PartSize, oss.Routines(info.Routines), oss.Checkpoint(info.Checkpoint, info.CheckpointFile))
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_part_upload_failed}")
	}

	return true, nil
}

// MultipartPartUpload 分片上传，将要上传的较大文件（Object）分成多个分片（Part）来分别上传
func (s *sOssAliyun) MultipartPartUpload(ctx context.Context, info *oss_model.MultipartPartPutObject) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// 将本地文件分片，且分片数量指定为3。
	chunks, err := oss.SplitFileByPartNum(info.FilePath, info.Chunks)
	fd, err := os.Open(info.FilePath)
	defer fd.Close()

	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(info.Expires),
		// 指定该Object被下载时的网页缓存行为。
		oss.CacheControl(info.CacheControl),
		// 指定该Object的内容编码格式。
		oss.ContentEncoding(info.ContentEncoding),

		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}

	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。 UploadId: 59ECF5FAAC8A496E8D6C22E611F08BEF
	imur, err := bucket.InitiateMultipartUpload(info.ObjectKey, options...)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_multipart_upload_failed}")
	}
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			return false, fmt.Errorf("{#error_oss_aliyun_multipart_upload_failed}")
		}

		parts = append(parts, part)
	}

	// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 步骤3：完成分片上传，指定文件读写权限为公共读。
	cmur, err := bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_multipart_upload_failed}")
	}
	fmt.Println("cmur:", cmur)

	return true, nil
}

// DownloadFile 下载文件到本地
func (s *sOssAliyun) DownloadFile(ctx context.Context, info *oss_model.DownLoadFile) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	options := make([]oss.Option, 0)

	// 如果是下载压缩过后的图片，需要设置压缩格式
	if info.ContentEncoding != "" {
		options = []oss.Option{
			oss.ContentEncoding(info.ContentEncoding),
			oss.SetHeader("Accept-Encoding", info.ContentEncoding),
		}
	}

	// 下载文件
	err = bucket.GetObjectToFile(info.ObjectKey, info.FilePath, options...)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_download_failed}")
	}

	return err == nil, err
}

// PartPutDownload 断点续传下载
func (s *sOssAliyun) PartPutDownload(ctx context.Context, info *oss_model.PartPutObject) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// 断点续传下载文件(文件key，源文件，分片大小，并发数，是否开启断点续传，断点续传文件)
	err = bucket.DownloadFile(info.ObjectKey, info.FilePath, info.PartSize, oss.Routines(info.Routines), oss.Checkpoint(info.Checkpoint, info.CheckpointFile))
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_part_download_failed}")
	}

	return err == nil, err
}

// GetObject 获取文件
func (s *sOssAliyun) GetObject(ctx context.Context, info *oss_model.GetFile) ([]byte, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return nil, err
	}

	// 获取文件
	object, err := bucket.GetObject(info.BucketName)
	if err != nil {
		return nil, fmt.Errorf("{#error_oss_aliyun_get_file_failed}")
	}
	defer object.Close()

	// 读取文件内容
	buf := make([]byte, 0)
	n, err := object.Read(buf)
	if err != nil || n == 0 {
		return nil, fmt.Errorf("{#error_oss_aliyun_read_file_failed}")
	}

	fmt.Println(string(buf))

	return buf, err
}

// GetFile 上传并获取文件 (OSS仅支持查询CSV文件和JSON文件)
func (s *sOssAliyun) GetFile(ctx context.Context, info *oss_model.GetFile) ([]byte, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return nil, err
	}

	// 查询CSV文件 | 查询JSON文件
	selReq := oss.SelectRequest{}
	selReq.Expression = `select * from ossObject`
	body, err := bucket.SelectObject(info.ObjectKey, selReq)

	if err != nil {
		return nil, fmt.Errorf("{#error_oss_aliyun_query_file_failed}")
	}

	// 读取内容。
	fileInfo, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, fmt.Errorf("{#error_oss_aliyun_read_file_failed}")
	}

	defer body.Close()

	fmt.Println(string(fileInfo))

	return fileInfo, err
}

// DeleteFile 删除文件
func (s *sOssAliyun) DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// 删除单个文件。
	err = bucket.DeleteObject(info.ObjectKey)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_delete_file_failed}")
	}

	return err == nil, err
}

// DeleteFileList 批量删除文件
func (s *sOssAliyun) DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// 删除文件 oss.DeleteObjectsQuiet(true)表示不返回删除结果
	_, err = bucket.DeleteObjects(info.ObjectKeys)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_delete_files_failed}")
	}

	return err == nil, err
}

// QueryFiles 列举文件  列举指定存储空间（Bucket）下的所有文件（Object）、指定前缀的文件、指定目录下的文件和子目录("/")
func (s *sOssAliyun) QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return nil, err
	}

	// 列举所有文件。
	continueToken := ""

	// 默认最大100个   注意： 在OSS里面没有什么文件夹可分，文件夹也看作是文件
	var oInfoList []oss_model.ObjectInfoRes

	options := []oss.Option{
		oss.ContinuationToken(continueToken),
		oss.Prefix(info.Prefix),
		oss.Delimiter(info.Delimiter),
		oss.StartAfter(info.StartAfter),
		oss.MaxKeys(info.MaxKeys),
	}

	// 列举文件  指定文件夹：prefix(文件夹名称) Delimiter(不填)
	lsRes, err := bucket.ListObjectsV2(options...)
	if err != nil {
		return nil, fmt.Errorf("{#error_oss_aliyun_list_files_failed}")
	}

	for _, object := range lsRes.Objects {
		var oInfo oss_model.ObjectInfoRes

		oInfo.ObjectKey = object.Key
		oInfo.ObjectType = object.Type
		oInfo.ObjectSize = object.Size
		oInfo.ObjectETag = object.ETag
		oInfo.ObjectLastModified = object.LastModified
		oInfo.ObjectStorageClass = object.StorageClass
		oInfo.ObjectOwnerID = object.Owner.ID
		oInfo.ObjectOwnerDisplayName = object.Owner.DisplayName

		oInfoList = append(oInfoList, oInfo)
	}

	return oInfoList, err
}

// CopyFileToPath 将指定文件拷贝到指定位置
func (s *sOssAliyun) CopyFileToPath(ctx context.Context, info *oss_model.CopyFileToPath) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	/*
		example 使用示例：
		1、查找某一个存储空间中的文件List
		res, err := s.QueryFiles(ctx, info)
		2、拼接目标文件名
		for _, re := range res {
			objectName := re.ObjectKey  // 原始文件名
			destObjectName := "resource/" + objectName // 目标文件名

			3、将指定文件拷贝到指定位置
			_, err = bucket.CopyObject(objectName, destObjectName, options...)
		}
	*/

	//指定目标文件的元数据。
	expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	tag1 := oss.Tag{
		Key:   "a",
		Value: "1",
	}

	taggingInfo := oss.Tagging{
		Tags: []oss.Tag{tag1},
	}

	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		oss.Expires(expires),
		oss.SetTagging(taggingInfo),
		// 指定复制源Object的对象标签到目标 Object。
		oss.TaggingDirective(oss.TaggingCopy),
		// 指定创建目标Object时的访问权限ACL为私有。
		// oss.ObjectACL(oss.ACLPrivate),
		// 指定KMS托管的用户主密钥，该参数仅在x-oss-server-side-encryption为KMS时有效。
		//oss.ServerSideEncryptionKeyID("9468da86-3509-4f8d-a61e-6eab1eac****"),
		// 指定OSS创建目标Object时使用的服务器端加密算法。
		// oss.ServerSideEncryption("AES256"),
		// 指定复制源Object的元数据到目标Object。
		//oss.MetadataDirective(oss.MetaCopy),
		// 指定CopyObject操作时是否覆盖同名目标Object。此处设置为true，表示禁止覆盖同名Object。
		// oss.ForbidOverWrite(true),
		// 如果源Object的ETag值和您提供的ETag相等，则执行拷贝操作，并返回200 OK。
		//oss.CopySourceIfMatch("5B3C1A2E053D763E1B002CC607C5****"),
		// 如果源Object的ETag值和您提供的ETag不相等，则执行拷贝操作，并返回200 OK。
		//oss.CopySourceIfNoneMatch("5B3C1A2E053D763E1B002CC607C5****"),
		// 如果指定的时间早于文件实际修改时间，则正常拷贝文件，并返回200 OK。
		//oss.CopySourceIfModifiedSince(2021-12-09T07:01:56.000Z),
		// 如果指定的时间等于或者晚于文件实际修改时间，则正常拷贝文件，并返回200 OK。
		//oss.CopySourceIfUnmodifiedSince(2021-12-09T07:01:56.000Z),
		// 指定Object的存储类型。此处设置为Standard，表示标准存储类型。
		//oss.StorageClass("Standard"),
	}

	// 使用指定的元数据覆盖源文件的元数据。
	_, err = bucket.CopyObject(info.ObjectName, info.DestObjectName, options...)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_copy_file_failed}")
	}

	return false, nil
}

// createOssClient 创建oss客户端
func (s *sOssAliyun) createOssClient(endpoint, accessKeyId, accessKeySecret string) (*oss.Client, error) {
	// 查看版本
	fmt.Println("OSS Go SDK Version: ", oss.Version)
	// 创建oss客户端
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	//s.OssClient = client
	return client, nil
}

// CreateBucket 创建Bucket存储空间操作对象
func (s *sOssAliyun) CreateBucket(ctx context.Context, info oss_model.MustInfo) (*oss.Bucket, error) {
	// 通过BucketId 找到配置信息
	providerInfo := oss_entity.OssServiceProviderConfig{}
	err := s.dao.OssServiceProviderConfig.Ctx(ctx).Where(oss_do.OssServiceProviderConfig{Id: info.ProviderId}).Scan(&providerInfo)
	if err != nil {
		return nil, fmt.Errorf("{#error_oss_aliyun_provider_id_error}")
	}

	// 创建oss客户端
	client, _ := s.createOssClient(providerInfo.Endpoint, providerInfo.AccessKeyId, providerInfo.AccessKeySecret)

	// 创建存储空间。
	bucket, err := client.Bucket(info.BucketName)
	if err != nil {
		return nil, fmt.Errorf("{#error_oss_aliyun_create_bucket_failed}")
	}

	return bucket, nil
}

// GetObjectToFileWithURL 根据URL获取存储对象
func (s *sOssAliyun) GetObjectToFileWithURL(ctx context.Context, info oss_model.GetObjectToFileWithURL) (bool, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return false, err
	}

	// 填写步骤1获取的签名URL。
	signedURL := info.SingUrl
	// 使用签名URL将OSS文件下载到本地文件。
	err = bucket.GetObjectToFileWithURL(signedURL, info.FilePath)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_aliyun_download_by_url_failed}")
	}

	return err == nil, err
}

// GetFileSingURL 获取文件的签名访问URL
func (s *sOssAliyun) GetFileSingURL(ctx context.Context, info *oss_model.GetFileSingURL, styleStr ...string) (string, error) {
	// 创建存储空间Bucket
	bucket, err := s.CreateBucket(ctx, info.MustInfo)
	if err != nil {
		return "", err
	}

	// 填写文件完整路径，例如exampledir/exampleobject.txt。文件完整路径中不能包含Bucket名称。
	objectName := info.ObjectKey

	// x-oss-process: image/quality,q_90/format,jpg
	style := "image"
	if len(styleStr) > 0 {
		style += styleStr[0]
	}

	// 生成用于下载的签名URL，并指定签名URL的有效时间为60秒。
	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, info.ExpiredInSec, oss.Process(style))
	if err != nil {
		return "", fmt.Errorf("{#error_oss_aliyun_get_sign_url_failed}")
	}
	fmt.Printf("Sign Url:%s\n", signedURL)

	return signedURL, nil
}
