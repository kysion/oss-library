package oss

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/kconv"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
	"github.com/kysion/oss-library/oss_model/oss_do"
	"github.com/kysion/oss-library/oss_model/oss_entity"
	"github.com/yitter/idgenerator-go/idgen"
)

// Oss应用配置管理
type sAppConfig struct {
	modules oss_interface.IModules
	dao     *oss_dao.XDao
}

func NewOssAppConfig(modules oss_interface.IModules) oss_interface.IAppConfig {
	return &sAppConfig{
		modules: modules,
		dao:     modules.Dao(),
	}
}

// GetAppConfigByName 根据应用名称获取AppConfig
func (s *sAppConfig) GetAppConfigByName(ctx context.Context, appName string) (*oss_model.OssAppConfig, error) {
	if appName == "" {
		table := s.modules.Dao().OssAppConfig.Table()
		return nil, errors.New("应用名称不能为空" + table)
	}

	data := oss_entity.OssAppConfig{}

	err := s.modules.Dao().OssAppConfig.Ctx(ctx).Where(oss_do.OssAppConfig{AppName: appName}).Scan(&data)
	if err != nil {
		return nil, errors.New("根据应用名称获取应用信息失败" + err.Error() + s.dao.OssAppConfig.Table())
	}

	res := kconv.Struct[*oss_model.OssAppConfig](data, &oss_model.OssAppConfig{})

	return res, nil
}

// GetAppConfigById 根据id获取AppConfig
func (s *sAppConfig) GetAppConfigById(ctx context.Context, id int64) (*oss_model.OssAppConfig, error) {
	if id == 0 {
		return nil, errors.New("应用id不能为空" + s.dao.OssAppConfig.Table())

	}

	data := oss_entity.OssAppConfig{}

	err := s.dao.OssAppConfig.Ctx(ctx).Where(oss_do.OssAppConfig{Id: id}).Scan(&data)

	if err != nil {
		return nil, errors.New("根据应用id获取应用信息失败" + err.Error() + s.dao.OssAppConfig.Table())
	}

	res := kconv.Struct[*oss_model.OssAppConfig](data, &oss_model.OssAppConfig{})

	return res, nil
}

// GetAppAvailableNumber 账户用量统计 (上下文, 应用id) (当前应用剩余短信数量)
func (s *sAppConfig) GetAppAvailableNumber(ctx context.Context, id int64) (int64, error) {
	if id == 0 {
		return 0, errors.New("应用id不能为空" + s.dao.OssAppConfig.Table())
	}

	data := oss_entity.OssAppConfig{}

	err := s.dao.OssAppConfig.Ctx(ctx).Where(oss_do.OssAppConfig{Id: id}).Scan(&data)
	if err != nil {
		return 0, errors.New("根据应用id获取应用信息失败" + err.Error() + s.dao.OssAppConfig.Table())
	}

	// 注意：返回的是字节数，Bytes(字节) 40GB就=1024×1024×1024×40=42949672960Byte
	return data.AvailableNumber, nil
}

// CreateAppConfig 创建应用 (上下文, 应用编号, 花费数量)
func (s *sAppConfig) CreateAppConfig(ctx context.Context, config *oss_model.OssAppConfig) (bool, error) {
	// 应用名称查重
	count, _ := s.dao.OssAppConfig.Ctx(ctx).Where(oss_do.OssAppConfig{
		AppName: config.AppName,
	}).Count()

	if count > 0 {
		return false, errors.New("应用名称重复" + s.dao.OssAppConfig.Table())
	}

	// 生成id
	appConfig := oss_do.OssAppConfig{}
	gconv.Struct(config, &appConfig)
	appConfig.Id = idgen.NextId()
	appConfig.Status = 1 // 默认正常
	appConfig.UnionMainId = config.UnionMainId

	_, err := s.dao.OssAppConfig.Ctx(ctx).Insert(appConfig)
	if err != nil {
		return false, errors.New("应用创建失败" + s.dao.OssAppConfig.Table())
	}

	return true, nil
}

// UpdateAppNumber 更新应用使用数量 (上下文, 应用编号, 花费数量)
func (s *sAppConfig) UpdateAppNumber(ctx context.Context, id int64, fee uint64) (bool, error) {
	if id == 0 {
		return false, errors.New("应用id不能为空" + s.dao.OssAppConfig.Table())
	}

	// 获取原来的数量
	appConfig, err := s.GetAppConfigById(ctx, id)
	if err != nil {
		return false, err
	}

	newUseNum := appConfig.UseNumber + gconv.Int64(fee)
	newAvailableNum := appConfig.AvailableNumber - gconv.Int64(fee)

	affected, err := daoctl.UpdateWithError(s.dao.OssAppConfig.Ctx(ctx).
		Data(oss_do.OssAppConfig{UseNumber: newUseNum, AvailableNumber: newAvailableNum}).
		Where(oss_do.OssAppConfig{Id: id}))

	return affected > 0, nil
}
