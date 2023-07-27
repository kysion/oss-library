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
		DeleteFile(ctx context.Context, info oss_model.DeleteFile) (bool, error)
		DeleteFileList(ctx context.Context, info oss_model.DeleteFileList) (bool, error)
		QueryFiles(ctx context.Context, info *oss_model.QueryFileListReq) (res []oss_model.ObjectInfoRes, err error)
		CreateBucket(ctx context.Context, info oss_model.MustInfo) (*oss.Bucket, error)
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
		GetAppConfigByName(ctx context.Context, appName string) (*oss_model.OssAppConfig, error)
		GetAppConfigById(ctx context.Context, id int64) (*oss_model.OssAppConfig, error)
		GetAppAvailableNumber(ctx context.Context, id int64) (int64, error)
		CreateAppConfig(ctx context.Context, config *oss_model.OssAppConfig) (bool, error)
		UpdateAppNumber(ctx context.Context, id int64, fee uint64) (bool, error)
	}
	IBucketConfig interface {
		GetBucketById(ctx context.Context, id int64) (*oss_model.OssBucketConfig, error)
		CreateBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error)
	}
	IServiceProviderConfig interface {
		CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error)
		GetProviderById(ctx context.Context, id int64) (*oss_model.OssServiceProviderConfig, error)
		QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*oss_model.OssServiceProviderListRes, error)
		QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*oss_model.OssServiceProviderListRes, error)
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
