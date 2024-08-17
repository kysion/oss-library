package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/utility/base_permission"
)

type Permission struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	ParentId    int64       `json:"parentId"    orm:"parent_id"   description:"父级ID"`
	Name        string      `json:"name"        orm:"name"        description:"名称"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Identifier  string      `json:"identifier"  orm:"identifier"  description:"标识符"`
	Type        int         `json:"type"        orm:"type"        description:"类型：1api，2menu"`
	MatchMode   int         `json:"matchMode"   orm:"match_mode"  description:"匹配模式：ID：0，标识符：1"`
	IsShow      int         `json:"isShow"      orm:"is_show"     description:"是否显示：0不显示 1显示"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""`
}

type PermissionTree struct {
	*Permission
	Children []base_permission.IPermission `json:"children"       dc:"下级权限"`
}

func (d *PermissionTree) GetIsEqual(father base_permission.IPermission, childId base_permission.IPermission) bool {
	return father.GetId() == childId.GetParentId()
}
func (d *PermissionTree) SetChild(father base_permission.IPermission, branchArr []base_permission.IPermission) {
	father.SetItems(branchArr)
}
func (d *PermissionTree) RetFather(father base_permission.IPermission) bool {
	// 顶级的ParentId这块可以看一下保存的时候ParentId 默认值是多少
	return father.GetParentId() == 0
}

// 实现权限接口

func (d *PermissionTree) GetId() int64 {
	return d.Id

}
func (d *PermissionTree) GetParentId() int64 {
	return d.ParentId

}
func (d *PermissionTree) GetName() string {
	return d.Name

}
func (d *PermissionTree) GetDescription() string {
	return d.Description

}
func (d *PermissionTree) GetIdentifier() string {
	return d.Identifier

}
func (d *PermissionTree) GetType() int {
	return d.Type

}
func (d *PermissionTree) GetMatchMode() int {
	return d.MatchMode

}
func (d *PermissionTree) GetIsShow() int {
	return d.IsShow

}
func (d *PermissionTree) GetSort() int {
	return d.Sort

}
func (d *PermissionTree) GetItems() []base_permission.IPermission {
	return d.Children

}
func (d *PermissionTree) GetData() interface{} {
	return d
}

func (d *PermissionTree) SetId(val int64) base_permission.IPermission {
	d.Id = val
	return d
}
func (d *PermissionTree) SetParentId(val int64) base_permission.IPermission {
	d.ParentId = val
	return d

}
func (d *PermissionTree) SetName(val string) base_permission.IPermission {
	d.Name = val
	return d

}
func (d *PermissionTree) SetDescription(val string) base_permission.IPermission {
	d.Description = val
	return d

}
func (d *PermissionTree) SetIdentifier(val string) base_permission.IPermission {
	d.Identifier = val
	return d

}
func (d *PermissionTree) SetType(val int) base_permission.IPermission {
	d.Type = val
	return d

}
func (d *PermissionTree) SetMatchMode(val int) base_permission.IPermission {
	d.MatchMode = val
	return d

}
func (d *PermissionTree) SetIsShow(val int) base_permission.IPermission {
	d.IsShow = val
	return d

}
func (d *PermissionTree) SetSort(val int) base_permission.IPermission {
	d.Sort = val
	return d

}
func (d *PermissionTree) SetItems(val []base_permission.IPermission) base_permission.IPermission {
	d.Children = val
	return d
}
