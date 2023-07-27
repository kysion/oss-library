package oss_v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/oss-library/api/oss_api"
)

type GetAppConfigByIdReq struct {
	g.Meta ` method:"post" summary:"根据Id查询应用" tags:"Oss配置"`

	oss_api.GetAppConfigByIdReq
}

type CreateAppConfigReq struct {
	g.Meta ` method:"post" summary:"创建应用" tags:"Oss配置"`

	oss_api.CreateAppConfigReq
}

type GetAppAvailableNumberReq struct {
	g.Meta ` method:"post" summary:"账户用量统计" tags:"Oss配置"`

	oss_api.GetAppAvailableNumberReq
}

type RegisterBucketReq struct {
	g.Meta ` method:"post" summary:"创建存储空间" tags:"Oss配置"`

	oss_api.RegisterBucketReq
}

type GetBucketByIdReq struct {
	g.Meta ` method:"post" summary:"根据id查询存储空间|信息" tags:"Oss配置"`

	oss_api.GetBucketByIdReq
}

type CreateProviderReq struct {
	g.Meta ` method:"post" summary:"添加渠道商" tags:"Oss配置"`

	oss_api.CreateProviderReq
}

type QueryProviderByNoReq struct {
	g.Meta ` method:"post" summary:"根据No编号获取渠道商|列表" tags:"Oss配置"`

	oss_api.QueryProviderByNoReq
}

type GetProviderByIdReq struct {
	g.Meta ` method:"post" summary:"根据id查询渠道商配置|信息" tags:"Oss配置"`

	oss_api.GetProviderByIdReq
}

type QueryProviderListReq struct {
	g.Meta ` method:"post" summary:"获取渠道商列表" tags:"Oss配置"`

	oss_api.QueryProviderListReq
}
