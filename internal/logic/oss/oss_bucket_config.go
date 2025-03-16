package oss

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
	"github.com/kysion/oss-library/oss_model/oss_do"
	"github.com/kysion/oss-library/oss_model/oss_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// Bucket 存储空间管理
type sBucketConfig struct {
	modules oss_interface.IModules
	dao     *oss_dao.XDao
}

func NewBucketConfig(modules oss_interface.IModules) oss_interface.IBucketConfig {
	return &sBucketConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetBucketById 根据id获取Bucket配置信息
func (s *sBucketConfig) GetBucketById(ctx context.Context, id int64) (*oss_model.OssBucketConfig, error) {
	if id == 0 {
		return nil, fmt.Errorf("{#error_oss_bucket_config_id_empty}")
	}

	data := oss_entity.OssBucketConfig{}

	err := s.dao.OssBucketConfig.Ctx(ctx).Where(oss_do.OssBucketConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, fmt.Errorf("{#error_oss_bucket_config_get_by_id_failed}")
	}

	res := kconv.Struct(data, &oss_model.OssBucketConfig{})

	return res, nil
}

// CreateBucket 创建存储空间配置信息 (上下文, 存储空间信息)
func (s *sBucketConfig) CreateBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error) {
	// 判断同一个渠道商下面是否名称重复
	count, err := s.dao.OssBucketConfig.Ctx(ctx).Where(oss_do.OssBucketConfig{
		BucketName: info.BucketName,
		ProviderNo: info.ProviderNo,
	}).Count()

	if err != nil {
		return false, fmt.Errorf("{#error_oss_bucket_config_check_name_failed}")
	}

	if count > 0 {
		return false, fmt.Errorf("{#error_oss_bucket_config_name_duplicate}")
	}

	// 生成id
	data := oss_do.OssBucketConfig{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.State = 1 // 默认正常
	// data.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	_, err = s.dao.OssBucketConfig.Ctx(ctx).Insert(data)
	if err != nil {
		return false, fmt.Errorf("{#error_oss_bucket_config_create_failed}")
	}

	return true, nil
}

// GetByBucketNameAndProviderNo 根据渠道商编号和Bucket存储对象名称获取存储对象
func (s *sBucketConfig) GetByBucketNameAndProviderNo(ctx context.Context, bucketName, providerNo string, state ...int) (*oss_model.OssBucketConfig, error) {
	if bucketName == "" {
		return nil, fmt.Errorf("{#error_oss_bucket_config_name_empty}")
	}

	data := oss_entity.OssBucketConfig{}
	model := s.dao.OssBucketConfig.Ctx(ctx).Where(oss_do.OssBucketConfig{BucketName: bucketName, ProviderNo: providerNo})
	if len(state) > 0 {
		model = model.Where(oss_do.OssBucketConfig{State: state})
	}

	err := model.Scan(&data)
	if err != nil {
		return nil, fmt.Errorf("{#error_oss_bucket_config_get_by_name_provider_failed}")
	}

	res := kconv.Struct(data, &oss_model.OssBucketConfig{})

	return res, nil
}
