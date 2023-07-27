package osser

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/oss-library/oss_model"
)

// 抽象接口定义
//  - OSS渠道管理
// 	- OSS应用资源管理
//  - 创建存储空间
//  - 上传文件
//	- 下载文件
//	- 列举文件
//	- 删除文件

type OSSer interface {
	// ---------------------------------------- 渠道商管理 ---------------------------------------------------

	// CreateProvider 添加渠道商
	CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error)

	// QueryProviderByNo 根据No编号获取渠道商列表
	QueryProviderByNo(ctx context.Context, no string, params base_model.SearchParams) (*oss_model.OssServiceProviderListRes, error)

	// QueryProviderList 获取渠道商列表
	QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*oss_model.OssServiceProviderListRes, error)

	// GetProviderById 根据id获取渠道商配置信息
	GetProviderById(ctx context.Context, id int64) (*oss_model.OssServiceProviderConfig, error)

	// --------------------------------------  应用管理 ----------------------------------------------------

	// GetAppConfigById 根据应用id查询应用
	GetAppConfigById(ctx context.Context, id int64) (*oss_model.OssAppConfig, error)

	// CreateAppConfig 创建应用 (上下文, 应用信息)
	CreateAppConfig(ctx context.Context, config *oss_model.OssAppConfig) (bool, error)

	// GetAppAvailableNumber 账户用量统计 (上下文, 应用Id) (当前应用剩余存空间字节数,err)
	GetAppAvailableNumber(ctx context.Context, appId int64) (int64, error)

	// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
	UpdateAppNumber(ctx context.Context, appId int64, fee uint64) (bool, error)

	// ---------------------------------------- 存储空间管理 ------------------------------------------------

	// CreateBucket 创建存储空间
	CreateBucket(ctx context.Context, info *oss_model.MustInfo) (*oss.Bucket, error)

	// GetBucketById 获取存储空间配置信息
	GetBucketById(ctx context.Context, id int64) (*oss_model.OssBucketConfig, error)

	// ---------------------------------------- 文件操作 ---------------------------------------------------

	// UploadFile 上传文件
	UploadFile(ctx context.Context, info *oss_model.PutObject) (bool, error)

	// PartPutUpload 断点续传文件
	PartPutUpload(ctx context.Context, info *oss_model.PartPutObject) (bool, error)

	// MultipartPartUpload 分片上传，将要上传的较大文件（Object）分成多个分片（Part）来分别上传
	MultipartPartUpload(ctx context.Context, info *oss_model.MultipartPartPutObject) (bool, error)

	// DownloadFile 下载文件到本地
	DownloadFile(ctx context.Context, file *oss_model.DownLoadFile) (bool, error)

	// PartPutDownload 断点续传下载
	PartPutDownload(ctx context.Context, info *oss_model.PartPutObject) (bool, error)

	// GetFile 获取文件
	GetFile(ctx context.Context, info *oss_model.GetFile) ([]byte, error)

	// DeleteFile 删除文件
	DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error)

	// DeleteFileList 批量删除文件
	DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error)

	// QueryFiles 列举文件  列举指定存储空间（Bucket）下的所有文件（Object）、指定前缀的文件、指定目录下的文件和子目录("/")
	QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error)
}
