// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package oss_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OssServiceProviderConfig is the golang structure for table oss_service_provider_config.
type OssServiceProviderConfig struct {
	Id              int64       `json:"id"              description:""`
	ProviderName    string      `json:"providerName"    description:"渠道商名称"`
	ProviderNo      string      `json:"providerNo"      description:"渠道商编号"`
	AccessKeyId     string      `json:"accessKeyId"     description:"身份标识"`
	AccessKeySecret string      `json:"accessKeySecret" description:"身份认证密钥"`
	Token           string      `json:"token"           description:"安全令牌"`
	BasePath        string      `json:"basePath"        description:"域名"`
	Endpoint        string      `json:"endpoint"        description:"bucket调用域名"`
	Remark          string      `json:"remark"          description:"备注"`
	Status          int         `json:"status"          description:"状态：0禁用 1启用"`
	ExtJson         string      `json:"extJson"         description:"拓展字段"`
	Region          string      `json:"region"          description:"地域"`
	CreatedAt       *gtime.Time `json:"createdAt"       description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       description:""`
	DeletedAt       *gtime.Time `json:"deletedAt"       description:""`
}