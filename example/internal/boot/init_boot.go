package boot

import (
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/oss-library/example/model"
)

func init() {
	//oss_global.Global.Modules.SetI18n(nil)
	//oss_global.Global.PermissionTree = boot.InitPermission(consts.Global.Modules)

	base_permission.InitializePermissionFactory(func() base_permission.IPermission {
		return &model.PermissionTree{
			Permission: &model.Permission{},
			Children:   nil,
		}
	})
}
