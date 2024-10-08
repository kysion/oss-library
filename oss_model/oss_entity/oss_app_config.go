// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package oss_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OssAppConfig is the golang structure for table oss_app_config.
type OssAppConfig struct {
	Id              int64       `json:"id"              orm:"id"               description:""`
	AppName         string      `json:"appName"         orm:"app_name"         description:"应用名称"`
	AvailableNumber int64       `json:"availableNumber" orm:"available_number" description:"可用总量"`
	CurrentLimiting int64       `json:"currentLimiting" orm:"current_limiting" description:"总量"`
	UseNumber       int64       `json:"useNumber"       orm:"use_number"       description:"使用量"`
	StartTime       *gtime.Time `json:"startTime"       orm:"start_time"       description:"生效时间"`
	EndTime         *gtime.Time `json:"endTime"         orm:"end_time"         description:"失效时间"`
	Status          int         `json:"status"          orm:"status"           description:"状态：0禁用 1正常"`
	Remark          string      `json:"remark"          orm:"remark"           description:"备注"`
	UnionMainId     int64       `json:"unionMainId"     orm:"union_main_id"    description:"主体id"`
	BucketId        int64       `json:"bucketId"        orm:"bucket_id"        description:"存储空间id"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:""`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"       description:""`
}
