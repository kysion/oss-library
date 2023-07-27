package boot

import (
	"context"
	"github.com/kysion/base-library/utility/base_permission"
	"github.com/kysion/oss-library/oss_interface"
)

// InitPermission 初始化权限树
func InitPermission(module oss_interface.IModules, permission base_permission.IPermission) []base_permission.IPermission {
	result := []base_permission.IPermission{
		// OSS
		permission.
			SetId(5947986066667901).
			SetName(module.T(context.TODO(), "{#OssName}")).
			SetIdentifier(module.GetConfig().Identifier.Oss).
			SetType(1).
			SetIsShow(1).
			SetItems([]base_permission.IPermission{}),
	}
	return result
}
