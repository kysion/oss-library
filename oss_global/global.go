package oss_global

import (
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
	"github.com/kysion/oss-library/oss_modules"
)

type global struct {
	Modules oss_interface.IModules
}

var (
	Global = global{
		Modules: oss_modules.NewModules(
			&oss_model.Config{
				HardDeleteWaitAt: 0,
				KeyIndex:         "Oss",
				RoutePrefix:      "/oss",
				I18nName:         "oss",
				StoragePath:      "./resource/oss",
				Identifier: oss_model.Identifier{
					Oss:                      "oss",
					OssAppConfig:             "ossAppConfig",
					OssServiceProviderConfig: "ossServiceProviderConfig",
					OssBucketConfig:          "ossBucketConfig",
				},
			},
			&oss_dao.XDao{ // 以下为业务层实例化dao模型，如果不是使用默认模型时需要将自定义dao模型作为参数传进去
				OssAppConfig:             oss_dao.NewOssAppConfig(),
				OssServiceProviderConfig: oss_dao.NewOssServiceProviderConfig(),
				OssBucketConfig:          oss_dao.NewOssBucketConfig(),
			},
		),
	}
)
