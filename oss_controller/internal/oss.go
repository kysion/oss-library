package internal

import (
	"context"
	"fmt"

	"github.com/kysion/oss-library/api/oss_api"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_interface/i_controller"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_enum"
)

// OssController 部分接口需要根据不同的请求调用不同的存储平台
type OssController struct {
	i_controller.IOss
	modules oss_interface.IModules
}

// Oss 存储服务
var Oss = func(modules oss_interface.IModules) i_controller.IOss {
	return &OssController{
		modules: modules,
	}
}

// GetAppConfigById 根据应用id查询应用
func (c *OssController) GetAppConfigById(ctx context.Context, req *oss_api.GetAppConfigByIdReq) (*oss_api.OssAppConfigRes, error) {
	ret, err := c.modules.OssAppConfig().GetAppConfigById(ctx, req.Id)

	return (*oss_api.OssAppConfigRes)(ret), err
}

// CreateAppConfig 创建应用 (上下文, 应用编号, 花费数量)
func (c *OssController) CreateAppConfig(ctx context.Context, req *oss_api.CreateAppConfigReq) (oss_api.BoolRes, error) {
	ret, err := c.modules.OssAppConfig().CreateAppConfig(ctx, &req.OssAppConfig)

	return ret == true, err
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用编号) (当前应用剩余存储空间余量)
func (c *OssController) GetAppAvailableNumber(ctx context.Context, req *oss_api.GetAppAvailableNumberReq) (oss_api.IntRes, error) {
	ret, err := c.modules.OssAppConfig().GetAppAvailableNumber(ctx, req.Id)

	return (oss_api.IntRes)(ret), err
}

// RegisterBucket 添加Bucket存储空间
func (c *OssController) RegisterBucket(ctx context.Context, req *oss_api.RegisterBucketReq) (res oss_api.BoolRes, err error) {
	var ret bool
	// 选择对应平台
	switch req.OssBucketConfig.ProviderNo {
	// 阿里云
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().RegisterBucket(ctx, &req.OssBucketConfig)
	// 腾讯云
	case oss_enum.Oss.Type.Tencent.Code():
		ret, err = c.modules.OssTencent().RegisterBucket(ctx, &req.OssBucketConfig)

	// 华为云
	case oss_enum.Oss.Type.Huawei.Code():

	// 七牛云
	case oss_enum.Oss.Type.Qiniu.Code():

	default:
		return false, fmt.Errorf("{#error_oss_controller_provider_not_supported}")
	}

	return ret == true, err
}

// GetBucketById 根据id获取Bucket配置信息
func (c *OssController) GetBucketById(ctx context.Context, req *oss_api.GetBucketByIdReq) (*oss_api.OssBucketConfigRes, error) {
	ret, err := c.modules.OssBucketConfig().GetBucketById(ctx, req.Id)

	return (*oss_api.OssBucketConfigRes)(ret), err
}

// CreateProvider 添加渠道商
func (c *OssController) CreateProvider(ctx context.Context, req *oss_api.CreateProviderReq) (res *oss_api.OssServiceProviderConfigRes, err error) {
	var ret *oss_model.OssServiceProviderConfig
	// 选择对应平台
	switch req.OssServiceProviderConfig.ProviderNo {
	// 阿里云
	case oss_enum.Oss.Type.Aliyun.Code():
		ret, err = c.modules.OssAliyun().CreateProvider(ctx, &req.OssServiceProviderConfig)

	// 腾讯云
	case oss_enum.Oss.Type.Tencent.Code():
		ret, err = c.modules.OssTencent().CreateProvider(ctx, &req.OssServiceProviderConfig)

	// 华为云
	case oss_enum.Oss.Type.Huawei.Code():

	// 七牛云
	case oss_enum.Oss.Type.Qiniu.Code():

	// 私有云

	// 本地云

	default:
		return nil, fmt.Errorf("{#error_oss_controller_provider_not_supported}")
	}

	// ret, err := c.modules.OssServiceProviderConfig().CreateProvider(ctx, &req.OssServiceProviderConfig)

	return (*oss_api.OssServiceProviderConfigRes)(ret), err
}

// GetProviderById 根据Id获取渠道商
func (c *OssController) GetProviderById(ctx context.Context, req *oss_api.GetProviderByIdReq) (*oss_api.OssServiceProviderConfigRes, error) {
	ret, err := c.modules.OssServiceProviderConfig().GetProviderById(ctx, req.Id)

	return (*oss_api.OssServiceProviderConfigRes)(ret), err
}

// QueryProviderByNo 根据No编号获取渠道商列表
func (c *OssController) QueryProviderByNo(ctx context.Context, req *oss_api.QueryProviderByNoReq) (*oss_api.OssServiceProviderConfigListRes, error) {
	ret, err := c.modules.OssServiceProviderConfig().QueryProviderByNo(ctx, req.No, &req.SearchParams)

	return (*oss_api.OssServiceProviderConfigListRes)(ret), err
}

// QueryProviderList 获取渠道商列表
func (c *OssController) QueryProviderList(ctx context.Context, req *oss_api.QueryProviderListReq) (*oss_api.OssServiceProviderConfigListRes, error) {
	ret, err := c.modules.OssServiceProviderConfig().QueryProviderList(ctx, &req.SearchParams, false)

	return (*oss_api.OssServiceProviderConfigListRes)(ret), err
}
