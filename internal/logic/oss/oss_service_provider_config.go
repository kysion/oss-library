package oss

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
	"github.com/kysion/oss-library/oss_model/oss_do"
	"github.com/kysion/oss-library/oss_model/oss_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// 渠道商管理
type sServiceProviderConfig struct {
	modules oss_interface.IModules
	dao     *oss_dao.XDao
}

func NewOssServiceProviderConfig(modules oss_interface.IModules) oss_interface.IServiceProviderConfig {
	return &sServiceProviderConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// CreateProvider 添加渠道商
func (s *sServiceProviderConfig) CreateProvider(ctx context.Context, info *oss_model.OssServiceProviderConfig) (*oss_model.OssServiceProviderConfig, error) {
	model := s.dao.OssServiceProviderConfig.Ctx(ctx)

	// 插入渠道商配置信息
	data := oss_do.OssServiceProviderConfig{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.CreatedAt = gtime.Now()
	// 渠道商默认是可用状态
	data.Status = 1
	data.ExtJson = nil

	_, err := model.OmitNilData().Insert(data)

	if err != nil {
		return nil, errors.New("渠道商添加失败" + s.dao.OssServiceProviderConfig.Table())
	}

	return s.GetProviderById(ctx, gconv.Int64(data.Id))
}

// GetProviderById 根据ID获取渠道商
func (s *sServiceProviderConfig) GetProviderById(ctx context.Context, id int64) (*oss_model.OssServiceProviderConfig, error) {
	if id == 0 {
		return nil, errors.New("渠道商id不能为空" + s.dao.OssServiceProviderConfig.Table())
	}

	data := oss_entity.OssServiceProviderConfig{}

	err := s.dao.OssServiceProviderConfig.Ctx(ctx).Where(oss_do.OssServiceProviderConfig{Id: id}).Scan(&data)
	if err != nil {
		return nil, errors.New("根据id获取渠道商信息失败：" + err.Error() + s.dao.OssServiceProviderConfig.Table())
	}

	res := kconv.Struct[*oss_model.OssServiceProviderConfig](data, &oss_model.OssServiceProviderConfig{})

	return res, nil
}

// QueryProviderByNo 根据No编号获取渠道商列表
func (s *sServiceProviderConfig) QueryProviderByNo(ctx context.Context, no string, params *base_model.SearchParams) (*oss_model.OssServiceProviderListRes, error) {
	if no == "" {
		return nil, errors.New("渠道商编号不能为空" + s.dao.OssServiceProviderConfig.Table())
	}

	res, err := daoctl.Query[oss_entity.OssServiceProviderConfig](s.dao.OssServiceProviderConfig.Ctx(ctx).Where(
		oss_do.OssServiceProviderConfig{ProviderNo: no}),
		params,
		false)

	if err != nil {
		return nil, errors.New("根据编号获取渠道商信息失败：" + err.Error() + s.dao.OssServiceProviderConfig.Table())
	}

	ret := kconv.Struct[*oss_model.OssServiceProviderListRes](res, &oss_model.OssServiceProviderListRes{})

	return ret, nil
}

// QueryProviderList 获取渠道商列表
func (s *sServiceProviderConfig) QueryProviderList(ctx context.Context, search *base_model.SearchParams, isExport bool) (*oss_model.OssServiceProviderListRes, error) {
	result, err := daoctl.Query[*oss_model.OssServiceProviderConfig](s.dao.OssServiceProviderConfig.Ctx(ctx), search, isExport)

	return (*oss_model.OssServiceProviderListRes)(result), err
}
