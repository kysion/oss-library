package oss_model

import (
	"encoding/xml"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"time"
)

// MustInfo 通过这两个去数据库查询出具体的oss配置信息
type MustInfo struct {
	ProviderId int64  `json:"providerId" dc:"渠道商Id" v:"required#渠道商id不能为空"`
	ProviderNo string `json:"providerNo"  dc:"渠道商编号" v:"required|in:local,aliyun,huawei,qiniu,minio,tencent#渠道商编号不能为空|渠道商校验失败"`
	BucketName string `json:"bucketName"  dc:"存储空间名称" v:"required#存储空间名称不能为空"`
}

// CreateBucket 创建存储空间请求对象
type CreateBucket struct {
	MustInfo
	AccessKeyId     string `json:"accessKeyId" dc:"身份标识" v:"required#身份标识不能为空"`
	AccessKeySecret string `json:"accessKeySecret" dc:"身份认证密钥" v:"required#身份认证密钥不能为空"`
	Endpoint        string `json:"endpoint" dc:"bucket调用域名" v:"required#bucket调用域名不能为空"`
}

// PutObject 文件上传请求对象
type PutObject struct {
	MustInfo
	Params    string `json:"params"  dc:"参数列表"`
	ObjectKey string `json:"objectKey" dc:"存储对象Name，也就是Key"`
	FilePath  string `json:"filePath" dc:"源文件存储路径"`
}

// PartPutObject 断点续传
type PartPutObject struct {
	MustInfo
	ObjectKey string `json:"objectKey" dc:"上传到OSS的文件名称，等同于objectName"`
	FilePath  string `json:"filePath" dc:"待上传的本地文件路径" `
	PartSize  int64  `json:"partSize" dc:"上传的分片大小，取值范围为100 KB~5 GB。默认值为100 KB。"`
	PartPutOptions
}
type PartPutOptions struct {
	Routines       int    `json:"routines" dc:"并发数，默认值是1，即不使用断点续传"`
	Checkpoint     bool   `json:"checkpoint" dc:"指定是否开启断点续传上传功能并设置Checkpoint文件。断点续传上传默认处于关闭状态。"`
	CheckpointFile string `json:"checkpointFile" dc:"本地文件名称，Checkpoint文件"`
}

// MultipartPartPutObject 分片上传
type MultipartPartPutObject struct {
	MustInfo
	ObjectKey string `json:"objectKey" dc:"上传到OSS的文件名称，等同于objectName"`
	FilePath  string `json:"filePath" dc:"待上传的本地文件路径" `
	Chunks    int    `json:"chunks" dc:"分片数量"`
	MultipartPartPutOptions
}
type MultipartPartPutOptions struct {
	Expires         time.Time `json:"expires" dc:"指定过期时间"`
	CacheControl    string    `json:"CacheControl" dc:"指定该Object被下载时的网页缓存行为，例如：no-cache"`
	ContentEncoding string    `json:"contentEncoding" dc:"指定object的内容编码格式，例如：gzip"`
	// 指定对返回的Key进行编码，目前支持URL编码。
	// 指定Object的存储类型。
}

// DownLoadFile 下载文件
type DownLoadFile struct {
	MustInfo
	ObjectKey       string `json:"objectKey" dc:"存储文件key"`
	FilePath        string `json:"filePath"  dc:"本地文件路径"`
	ContentEncoding string `json:"contentEncoding" dc:"压缩格式"`
}

// GetObjectToFileWithURL 根据路径下载文件
type GetObjectToFileWithURL struct {
	MustInfo
	FilePath string `json:"filePath"  dc:"本地存储的文件路径"`
	SingUrl  string `json:"singUrl" dc:"文件URL"`
}

// GetFileSingURL 根据路径下载文件
type GetFileSingURL struct {
	MustInfo
	ObjectKey    string `json:"objectKey" dc:"存储文件key"`
	ExpiredInSec int64  `json:"expiredInSec" dc:"URL过期时间，单位为秒"`
}

// GetFile 查询文件
type GetFile struct {
	MustInfo
	ObjectKey string `json:"objectKey" dc:"等同于objectName"`
	//LocalPath string `json:"localPath" dc:"本地文件路径"`
	Type string `json:"type" dc:"查询文件类型" v:"required|in:json,csv#查询类型不能为空|暂不支持此类型"`
	oss.SelectRequest
}

// DeleteFile 删除文件
type DeleteFile struct {
	MustInfo
	ObjectKey string `json:"objectKey" dc:"等同于objectName"`
}

// DeleteFileList 批量删除文件
type DeleteFileList struct {
	MustInfo
	ObjectKeys []string `json:"objectKey" dc:"等同于objectName"`
}

// QueryFileListReq 列举文件
type QueryFileListReq struct {
	MustInfo
	Prefix     string `json:"prefix" dc:"列举文件前缀"`
	MaxKeys    int    `json:"maxKeys" dc:"分页数量"`
	Delimiter  string `json:"delimiter" dc:"对文件名称进行分组的一个字符"`
	StartAfter string `json:"startAfter"  dc:"此次列举文件的起点"`
	// OSS没有文件夹的概念，所有元素都是以文件来存储。创建文件夹本质上来说是创建了一个大小为0并以正斜线（/）结尾的文件
	// 设置prefix为某个文件夹名称，则会列举以此prefix开头的文件
	// 设置了prefix的情况下，将delimiter设置为正斜线（/）
}

// CopyFileToPath 将指定文件拷贝到指定位置
type CopyFileToPath struct {
	MustInfo
	ObjectName     string `json:"objectName" dc:"原始文件名称"`
	DestObjectName string `json:"destObjectName" dc:"目标文件名称"`
}

// ObjectInfoRes 文件信息
type ObjectInfoRes struct {
	ObjectKey              string    `json:"objectKey" xml:"Key" dc:"存储文件key"`
	ObjectType             string    `json:"objectType" xml:"Type" dc:"文件类型"`
	ObjectSize             int64     `json:"objectSize" xml:"Size" dc:"文件大小"`
	ObjectETag             string    `json:"objectETag" xml:"ETag" dc:"文件标签"`
	ObjectLastModified     time.Time `json:"objectLastModified" xml:"LastModified" dc:"文件最后修改日期"`
	ObjectStorageClass     string    `json:"objectStorageClass" xml:"StorageClass" dc:"文件存储类型"`
	ObjectOwnerID          string    `json:"objectOwnerId" xml:"Owner" dc:"文件拥有者"`
	ObjectOwnerDisplayName string    `json:"objectOwnerDisplayName" xml:"objectOwnerDisplayName" dc:"文件显示名称"`
}
type ObjectInfoListRes []ObjectInfoRes

type Owner struct {
	XMLName     xml.Name `xml:"Owner"`
	ID          string   `xml:"ID"`          // Owner ID
	DisplayName string   `xml:"DisplayName"` // Owner's display name
}
