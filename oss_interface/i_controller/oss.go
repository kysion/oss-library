package i_controller

import (
	"context"
	"github.com/kysion/oss-library/api/oss_api"
)

type IOss interface {
	iModule

	GetAppConfigById(ctx context.Context, req *oss_api.GetAppConfigByIdReq) (*oss_api.OssAppConfigRes, error)

	CreateAppConfig(ctx context.Context, req *oss_api.CreateAppConfigReq) (oss_api.BoolRes, error)

	GetAppAvailableNumber(ctx context.Context, req *oss_api.GetAppAvailableNumberReq) (oss_api.IntRes, error)

	RegisterBucket(ctx context.Context, req *oss_api.RegisterBucketReq) (res oss_api.BoolRes, err error)

	GetBucketById(ctx context.Context, req *oss_api.GetBucketByIdReq) (*oss_api.OssBucketConfigRes, error)

	CreateProvider(ctx context.Context, req *oss_api.CreateProviderReq) (res *oss_api.OssServiceProviderConfigRes, err error)

	QueryProviderByNo(ctx context.Context, req *oss_api.QueryProviderByNoReq) (*oss_api.OssServiceProviderConfigListRes, error)

	GetProviderById(ctx context.Context, req *oss_api.GetProviderByIdReq) (*oss_api.OssServiceProviderConfigRes, error)

	QueryProviderList(ctx context.Context, req *oss_api.QueryProviderListReq) (*oss_api.OssServiceProviderConfigListRes, error)
}
