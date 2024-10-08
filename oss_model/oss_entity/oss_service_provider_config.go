// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package oss_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OssServiceProviderConfig is the golang structure for table oss_service_provider_config.
type OssServiceProviderConfig struct {
	Id              int64       `json:"id"              orm:"id"                description:""`
	ProviderName    string      `json:"providerName"    orm:"provider_name"     description:"渠道商名称"`
	ProviderNo      string      `json:"providerNo"      orm:"provider_no"       description:"渠道商编号"`
	AccessKeyId     string      `json:"accessKeyId"     orm:"access_key_id"     description:"身份标识"`
	AccessKeySecret string      `json:"accessKeySecret" orm:"access_key_secret" description:"身份认证密钥"`
	Token           string      `json:"token"           orm:"token"             description:"安全令牌"`
	Endpoint        string      `json:"endpoint"        orm:"endpoint"          description:"Bucket的地域域名"`
	Remark          string      `json:"remark"          orm:"remark"            description:"备注"`
	Status          int         `json:"status"          orm:"status"            description:"状态：0禁用 1启用"`
	ExtJson         string      `json:"extJson"         orm:"ext_json"          description:"拓展字段"`
	Region          string      `json:"region"          orm:"region"            description:"地域"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:""`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"        description:""`
	Priority        int         `json:"priority"        orm:"priority"          description:"优先级，使用默认选择优先级最高的"`
}
