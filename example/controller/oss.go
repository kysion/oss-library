package controller

import (
	"context"
	"github.com/kysion/oss-library/api/oss_api"
	"github.com/kysion/oss-library/api/oss_v1"
	"github.com/kysion/oss-library/oss_controller"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_interface/i_controller"
)

type OssController struct {
	i_controller.IOss
}

var Oss = func(modules oss_interface.IModules) *OssController {
	return &OssController{
		oss_controller.Oss(modules),
	}
}

func (c *OssController) GetModules() oss_interface.IModules {
	return c.IOss.GetModules()
}

func (c *OssController) GetAppConfigById(ctx context.Context, req *oss_v1.GetAppConfigByIdReq) (*oss_api.OssAppConfigRes, error) {
	return c.IOss.GetAppConfigById(ctx, &req.GetAppConfigByIdReq)
}

func (c *OssController) CreateAppConfig(ctx context.Context, req *oss_v1.CreateAppConfigReq) (oss_api.BoolRes, error) {
	return c.IOss.CreateAppConfig(ctx, &req.CreateAppConfigReq)
}

func (c *OssController) GetAppAvailableNumber(ctx context.Context, req *oss_v1.GetAppAvailableNumberReq) (oss_api.IntRes, error) {
	return c.IOss.GetAppAvailableNumber(ctx, &req.GetAppAvailableNumberReq)
}

func (c *OssController) RegisterBucket(ctx context.Context, req *oss_v1.RegisterBucketReq) (res oss_api.BoolRes, err error) {
	return c.IOss.RegisterBucket(ctx, &req.RegisterBucketReq)
}

func (c *OssController) GetBucketById(ctx context.Context, req *oss_v1.GetBucketByIdReq) (*oss_api.OssBucketConfigRes, error) {
	return c.IOss.GetBucketById(ctx, &req.GetBucketByIdReq)
}

func (c *OssController) CreateProvider(ctx context.Context, req *oss_v1.CreateProviderReq) (res *oss_api.OssServiceProviderConfigRes, err error) {
	return c.IOss.CreateProvider(ctx, &req.CreateProviderReq)
}

func (c *OssController) QueryProviderByNo(ctx context.Context, req *oss_v1.QueryProviderByNoReq) (*oss_api.OssServiceProviderConfigListRes, error) {
	return c.IOss.QueryProviderByNo(ctx, &req.QueryProviderByNoReq)
}

func (c *OssController) GetProviderById(ctx context.Context, req *oss_v1.GetProviderByIdReq) (*oss_api.OssServiceProviderConfigRes, error) {
	return c.IOss.GetProviderById(ctx, &req.GetProviderByIdReq)
}

func (c *OssController) QueryProviderList(ctx context.Context, req *oss_v1.QueryProviderListReq) (*oss_api.OssServiceProviderConfigListRes, error) {
	return c.IOss.QueryProviderList(ctx, &req.QueryProviderListReq)
}
