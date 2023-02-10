package form

import "gorm.io/gorm"

// CreateRoleRequest 创建菜单request
type CreateRoleRequest struct {
	gorm.Model
	Uid    string `json:"uid"`                      // 角色uid
	Name   string `json:"name" validate:"required"` // 角色名称
	Remark string `json:"remark"`                   // 角色描述
	Order  int    `json:"order"`                    // 排序
	Status int    `json:"status"`                   // 状态
}

// DelRoleRequest 删除菜单request
type DelRoleRequest struct {
	Uid string `json:"uid" validate:"required"` // uid
}

// UpdateRoleRequest 编辑菜单request
type UpdateRoleRequest struct {
	gorm.Model
	Uid    string `json:"uid" validate:"required"`  // 角色uid
	Name   string `json:"name" validate:"required"` // 角色名称
	Remark string `json:"remark"`                   // 角色描述
	Order  int    `json:"order"`                    // 排序
	Status int    `json:"status"`                   // 状态
}

// CurrentRoleAuthorizationMenuRequest 获取当前角色授权菜单request
type CurrentRoleAuthorizationMenuRequest struct {
	RoleUid string `json:"roleUid" validate:"required"` // roleUid
}

// SetCurrentRoleAuthorization 给当前角色授权
type SetCurrentRoleAuthorization struct {
	MenuUidList []string `json:"menuUidList" validate:"required"`
	RoleUid     string   `json:"roleUid" validate:"required"`
}
