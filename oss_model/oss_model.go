package oss_model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

// OssServiceProviderConfig 渠道商
type OssServiceProviderConfig struct {
	ProviderName    string `json:"providerName"    dc:"渠道商名称" v:"required#存储空间名称不能为空"`
	ProviderNo      string `json:"providerNo"      dc:"渠道商编号" v:"required|in:local,aliyun,huawei,qiniu,minio,tencent#渠道商编号不能为空|渠道商校验失败"`
	AccessKeyId     string `json:"accessKeyId"     dc:"身份标识"`
	AccessKeySecret string `json:"accessKeySecret" dc:"身份认证密钥"`
	Token           string `json:"token"           dc:"安全令牌"`
	BasePath        string `json:"basePath"        dc:"域名"`
	Endpoint        string `json:"endpoint"        dc:"bucket调用域名"`
	Remark          string `json:"remark"          dc:"备注"`
	Status          int    `json:"status"          dc:"状态：0禁用 1启用"`
	ExtJson         string `json:"extJson"         dc:"拓展字段"`
	Region          string `json:"region"          dc:"地域"`
}
type OssServiceProviderListRes base_model.CollectRes[*OssServiceProviderConfig]

// OssBucketConfig 存储空间
type OssBucketConfig struct {
	Id            int64  `json:"id" dc:"存储空间id"`
	BucketName    string `json:"bucketName"    dc:"存储空间名称"`
	Endpoint      string `json:"endpoint"      dc:"bucket调用域名"`
	StorageType   string `json:"storageType"   dc:"存储类型"`
	RedundantType string `json:"redundantType" dc:"冗余类型"`
	MonthlyFlow   int64  `json:"monthlyFlow"   dc:"当月流量"`
	VisitsNum     int    `json:"visitsNum"     dc:"访问次数"`
	UnionMainId   int64  `json:"unionMainId"   dc:"主体ID"`
	OwnerId       int64  `json:"ownerId"       dc:"拥有者ID"`
	ProviderNo    string `json:"providerNo"    dc:"渠道商编号" v:"required|in:local,aliyun,huawei,qiniu,minio,tencent#渠道商编号不能为空|渠道商校验失败"`
	State         int    `json:"state"         dc:"状态：0禁用 1正常"`
}
type OssBucketConfigListRes base_model.CollectRes[OssBucketConfig]

// OssAppConfig  应用资源配置
type OssAppConfig struct {
	AppName         string      `json:"appName"         dc:"应用名称"`
	AvailableNumber int64       `json:"availableNumber" dc:"可用总量"`
	CurrentLimiting int64       `json:"currentLimiting" dc:"总量"`
	UseNumber       int64       `json:"useNumber"       dc:"使用量"`
	StartTime       *gtime.Time `json:"startTime"       dc:"生效时间"`
	EndTime         *gtime.Time `json:"endTime"         dc:"失效时间"`
	Status          int         `json:"status"          dc:"状态：0禁用 1正常"`
	Remark          string      `json:"remark"          dc:"备注"`
	UnionMainId     int64       `json:"unionMainId"     dc:"主体id"`
	BucketId        int64       `json:"bucketId"        dc:"存储空间id"`
}
type OssAppConfigListRes base_model.CollectRes[OssAppConfig]
