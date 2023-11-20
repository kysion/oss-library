package oss_modules

import (
	"context"
	"github.com/kysion/oss-library/internal/logic/oss"

	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/kysion/oss-library/internal/boot"
	"github.com/kysion/oss-library/internal/logic/oss_aliyun"
	"github.com/kysion/oss-library/oss_consts"
	"github.com/kysion/oss-library/oss_interface"
	"github.com/kysion/oss-library/oss_model"
	"github.com/kysion/oss-library/oss_model/oss_dao"
)

type Modules struct {
	conf                     *oss_model.Config
	ossAppConfig             oss_interface.IAppConfig
	ossServiceProviderConfig oss_interface.IServiceProviderConfig
	ossBucketConfig          oss_interface.IBucketConfig

	ossAliyun  oss_interface.IOssAliyun
	ossTencent oss_interface.IOssTencent

	i18n *gi18n.Manager
	xDao *oss_dao.XDao
}

func (m *Modules) OssAppConfig() oss_interface.IAppConfig {
	return m.ossAppConfig
}

func (m *Modules) OssServiceProviderConfig() oss_interface.IServiceProviderConfig {
	return m.ossServiceProviderConfig
}

func (m *Modules) OssBucketConfig() oss_interface.IBucketConfig {
	return m.ossBucketConfig
}

func (m *Modules) OssAliyun() oss_interface.IOssAliyun {
	return m.ossAliyun
}

func (m *Modules) OssTencent() oss_interface.IOssTencent {
	return m.ossTencent
}

func (m *Modules) GetConfig() *oss_model.Config {
	return m.conf
}

func (m *Modules) T(ctx context.Context, content string) string {
	return m.i18n.Translate(ctx, content)
}

// Tf is alias of TranslateFormat for convenience.
func (m *Modules) Tf(ctx context.Context, format string, values ...interface{}) string {
	return m.i18n.TranslateFormat(ctx, format, values...)
}

func (m *Modules) SetI18n(i18n *gi18n.Manager) error {
	if i18n == nil {
		i18n = gi18n.New()
		i18n.SetLanguage("zh-CN")
		//err := i18n.SetPath("i18n/" + gstr.ToLower(m.conf.KeyIndex))
		err := i18n.SetPath("i18n/" + m.conf.I18nName)
		if err != nil {
			return err
		}
	}

	m.i18n = i18n
	return nil
}

func (m *Modules) Dao() *oss_dao.XDao {
	return m.xDao
}

func NewModules(
	conf *oss_model.Config,
	xDao *oss_dao.XDao,
) *Modules {
	module := &Modules{
		conf: conf,
		xDao: xDao,
	}

	// 初始化默认多语言对象
	module.SetI18n(nil)

	module.ossAppConfig = oss.NewOssAppConfig(module)
	module.ossServiceProviderConfig = oss.NewOssServiceProviderConfig(module)
	module.ossBucketConfig = oss.NewBucketConfig(module)
	module.ossAliyun = oss_aliyun.NewOssAliyun(module)

	// 权限树追加权限
	oss_consts.PermissionTree = append(oss_consts.PermissionTree, boot.InitPermission(module, nil)...)

	return module
}
