package oss

import (
	"context"
	"errors"
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
		return nil, errors.New("id不能为空" + s.dao.OssBucketConfig.Table())
	}

	data := oss_entity.OssBucketConfig{}

	err := s.dao.OssBucketConfig.Ctx(ctx).Where(oss_do.OssBucketConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("根据id获取存储空间信息失败" + err.Error() + s.dao.OssBucketConfig.Table())
	}

	res := kconv.Struct[*oss_model.OssBucketConfig](data, &oss_model.OssBucketConfig{})

	return res, nil
}

// CreateBucket 创建存储空间配置信息 (上下文, 存储空间信息)
func (s *sBucketConfig) CreateBucket(ctx context.Context, info *oss_model.OssBucketConfig) (bool, error) {
	// 判断同一个渠道商下面是否名称重复
	count, _ := s.dao.OssBucketConfig.Ctx(ctx).Where(oss_do.OssBucketConfig{
		BucketName: info.BucketName,
		ProviderNo: info.ProviderNo,
	}).Count()

	if count > 0 {
		return false, errors.New("存储空间名称重复" + s.dao.OssBucketConfig.Table())
	}

	// 生成id
	data := oss_do.OssBucketConfig{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.State = 1 // 默认正常
	// data.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	_, err := s.dao.OssBucketConfig.Ctx(ctx).Insert(data)
	if err != nil {
		return false, errors.New("存储空间创建失败" + s.dao.OssBucketConfig.Table())
	}

	return true, nil
}
