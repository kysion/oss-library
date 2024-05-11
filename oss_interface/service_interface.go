package oss_interface

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
)

type (
	IOssAliyun interface {
		CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error)
		RegisterBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error)
		UploadFile(ctx context.Context, info *oss_model.PutObject) (bool, error)
		PartPutUpload(ctx context.Context, info *oss_model.PartPutObject) (bool, error)
		MultipartPartUpload(ctx context.Context, info *oss_model.MultipartPartPutObject) (bool, error)
		DownloadFile(ctx context.Context, info *oss_model.DownLoadFile) (bool, error)
		PartPutDownload(ctx context.Context, info *oss_model.PartPutObject) (bool, error)
		GetFile(ctx context.Context, info *oss_model.GetFile) ([]byte, error)
		GetObject(ctx context.Context, info *oss_model.GetFile) ([]byte, error)
		DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error)
		DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error)
		QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error)
		CreateBucket(ctx context.Context, info oss_model.MustInfo) (*oss.Bucket, error)
		GetObjectToFileWithURL(ctx context.Context, info oss_model.GetObjectToFileWithURL) (bool, error)
		GetFileSingURL(ctx context.Context, info *oss_model.GetFileSingURL) (string, error)
	}

	IOssTencent interface {
		CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error)
		RegisterBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error)
		UploadFile(ctx context.Context, info *oss_model.PutObject) (bool, error)
		PartPutUpload(ctx context.Context, info *oss_model.PartPutObject) (bool, error)
		MultipartPartUpload(ctx context.Context, info *oss_model.MultipartPartPutObject) (bool, error)
		DownloadFile(ctx context.Context, info *oss_model.DownLoadFile) (bool, error)
		PartPutDownload(ctx context.Context, info *oss_model.PartPutObject) (bool, error)
		GetFile(ctx context.Context, info *oss_model.GetFile) ([]byte, error)
		DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error)
		DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error)
		QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error)
	}

	IAppConfig interface {
		// GetAppConfigByName 根据应用名称获取AppConfig
		GetAppConfigByName(ctx context.Context, appName string) (*oss_model.OssAppConfig, error)
		// GetAppConfigById 根据id获取AppConfig
		GetAppConfigById(ctx context.Context, id int64) (*oss_model.OssAppConfig, error)
		// GetAppAvailableNumber 账户用量统计 (上下文, 应用id) (当前应用剩余短信数量)
		GetAppAvailableNumber(ctx context.Context, id int64) (int64, error)
		// CreateAppConfig 创建应用 (上下文, 应用编号, 花费数量)
		CreateAppConfig(ctx context.Context, config *oss_model.OssAppConfig) (bool, error)
		// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
		UpdateAppNumber(ctx context.Context, id int64, fee uint64) (bool, error)
	}
	IBucketConfig interface {
		// GetBucketById 根据id获取Bucket配置信息
		GetBucketById(ctx context.Context, id int64) (*oss_model.OssBucketConfig, error)
		// CreateBucket 创建存储空间配置信息 (上下文, 存储空间信息)
		CreateBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error)
		// GetByBucketNameAndProviderNo 根据渠道商编号和Bucket存储对象名称获取存储对象
		GetByBucketNameAndProviderNo(ctx context.Context, bucketName, providerNo string, state ...int) (*oss_model.OssBucketConfig, error)
	}
	IServiceProviderConfig interface {
		// CreateProvider 添加渠道商
		CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error)
		// GetProviderById 根据ID获取渠道商
		GetProviderById(ctx context.Context, id int64) (*oss_model.OssServiceProviderConfig, error)
		// QueryProviderByNo 根据No编号获取渠道商列表
		QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*oss_model.OssServiceProviderListRes, error)
		// QueryProviderList 获取渠道商列表
		QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*oss_model.OssServiceProviderListRes, error)
		// GetProviderByPriority 根据优先级获取渠道商
		GetProviderByPriority(ctx context.Context, priority int) (*oss_model.OssServiceProviderConfig, error)
	}
)

type IModules interface {
	OssAppConfig() IAppConfig
	OssServiceProviderConfig() IServiceProviderConfig
	OssBucketConfig() IBucketConfig

	OssAliyun() IOssAliyun
	OssTencent() IOssTencent

	GetConfig() *oss_model.Config
	SetI18n(i18n *gi18n.Manager) error
	T(ctx context.Context, content string) string
	Tf(ctx context.Context, format string, values ...interface{}) string
	Dao() *oss_dao.XDao
}
