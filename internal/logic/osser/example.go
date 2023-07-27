package osser

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/oss-library/oss_model"
)

type Test struct {
	// 只要写了这个匿名字段，那么就可以实现接口的部分方法
	OSSer
}

// ---------------------------------------- 渠道商管理 ---------------------------------------------------

// CreateProvider 添加渠道商
func (s *Test) CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error) {
	return nil, nil
}

// QueryProviderByNo 根据No编号获取渠道商
func (s *Test) QueryProviderByNo(ctx context.Context, no string, params base_model.SearchParams) (*oss_model.OssServiceProviderListRes, error) {
	return nil, nil
}

// QueryProviderList 获取渠道商列表
func (s *Test) QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*oss_model.OssServiceProviderListRes, error) {
	return nil, nil
}

// GetProviderById 根据id获取渠道商配置信息
func (s *Test) GetProviderById(ctx context.Context, id int64) (*oss_model.OssServiceProviderConfig, error) {
	return nil, nil
}

// --------------------------------------  应用管理 ----------------------------------------------------

// GetAppConfigById 根据应用id查询应用
func (s *Test) GetAppConfigById(ctx context.Context, id int64) (*oss_model.OssAppConfig, error) {
	return nil, nil
}

// CreateAppConfig 创建应用 (上下文, 应用信息)
func (s *Test) CreateAppConfig(ctx context.Context, config *oss_model.OssAppConfig) (bool, error) {
	return false, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用Id) (当前应用剩余存空间字节数,err)
func (s *Test) GetAppAvailableNumber(ctx context.Context, appId int64) (int64, error) {
	return 0, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (s *Test) UpdateAppNumber(ctx context.Context, appId int64, fee uint64) (bool, error) {
	return false, nil
}

// ---------------------------------------- 存储空间管理 ------------------------------------------------

// CreateBucket 创建存储空间
func (s *Test) CreateBucket(ctx context.Context, info *oss_model.MustInfo) (*oss.Bucket, error) {
	return nil, nil
}

// GetBucketById 获取存储空间配置信息
func (s *Test) GetBucketById(ctx context.Context, id int64) (*oss_model.OssBucketConfig, error) {
	return nil, nil
}

// ---------------------------------------- 文件操作 ---------------------------------------------------

// UploadFile 上传文件
func (s *Test) UploadFile(ctx context.Context, info *oss_model.PutObject) (bool, error) {
	return false, nil
}

// PartPutUpload 断点续传文件
func (s *Test) PartPutUpload(ctx context.Context, info *oss_model.PartPutObject) (bool, error) {
	return false, nil
}

// MultipartPartUpload 分片上传，将要上传的较大文件（Object）分成多个分片（Part）来分别上传
func (s *Test) MultipartPartUpload(ctx context.Context, info *oss_model.MultipartPartPutObject) (bool, error) {
	return false, nil
}

// DownloadFile 下载文件到本地
func (s *Test) DownloadFile(ctx context.Context, info *oss_model.DownLoadFile) (bool, error) {
	return false, nil
}

// PartPutDownload 断点续传下载
func (s *Test) PartPutDownload(ctx context.Context, info *oss_model.PartPutObject) (bool, error) {
	return false, nil
}

// GetFile 获取文件
func (s *Test) GetFile(ctx context.Context, info *oss_model.GetFile) ([]byte, error) {
	return nil, nil
}

// DeleteFile 删除文件
func (s *Test) DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error) {
	return false, nil
}

// DeleteFileList 批量删除文件
func (s *Test) DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error) {
	return false, nil
}

// QueryFiles 列举文件  列举指定存储空间（Bucket）下的所有文件（Object）、指定前缀的文件、指定目录下的文件和子目录("/")
func (s *Test) QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error) {
	return nil, nil
}
