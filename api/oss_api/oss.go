package oss_api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/oss-library/oss_model"
)

type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type GetAppConfigByIdReq struct {
	Id int64 `json:"id" v:"required#应用id校验错误" dc:"应用id" `
}

type CreateAppConfigReq struct {
	oss_model.OssAppConfig
}

type GetAppAvailableNumberReq struct {
	Id int64 `json:"id" v:"required#应用id校验失败" dc:"应用ID"`
}

type RegisterBucketReq struct {
	oss_model.OssBucketConfig
}

type GetBucketByIdReq struct {
	Id int64 `json:"id" v:"required#存储空间id不能为空" dc:"存储空间id"`
}

type CreateProviderReq struct {
	oss_model.OssServiceProviderConfig
}

type QueryProviderByNoReq struct {
	No string `json:"no" v:"required#渠道商编号不能为空" dc:"渠道商编号"`
	base_model.SearchParams
}

type GetProviderByIdReq struct {
	Id int64 `json:"id" v:"required#渠道商id不能为空" dc:"渠道商id"`
}

type QueryProviderListReq struct {
	base_model.SearchParams
}

type OssAppConfigRes oss_model.OssAppConfig

type OssServiceProviderConfigRes oss_model.OssServiceProviderConfig

type OssServiceProviderConfigListRes oss_model.OssServiceProviderListRes

type OssBucketConfigRes oss_model.OssBucketConfig

type BoolRes bool
type StringRes string
type IntRes int
