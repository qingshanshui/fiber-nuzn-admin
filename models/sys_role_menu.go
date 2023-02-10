package models

import (
	"fiber-nuzn-api/initalize"
	"fiber-nuzn-api/pkg/utils"

	"gorm.io/gorm"
)

// SysRoleMenu 角色-菜单关系表
type SysRoleMenu struct {
	gorm.Model
	MenuUid string `json:"menuUid"`
	RoleUid string `json:"roleUid"`
}

// // TableName 重命名表
//
//	func (u *SysRoleMenu) TableName() string {
//		return "sys_role_menu"
//	}
//

// NewSysRoleMenu new一个空的结构体
func NewSysRoleMenu() *SysRoleMenu {
	return &SysRoleMenu{}
}

// Create 创建
func (u *SysRoleMenu) Create(menuUid, roleUid string) error {
	return initalize.DB.Exec("insert into sys_role_menu (created_at, updated_at, menu_uid, role_uid) VALUES (?,?,?,?)",
		utils.GetDayTime(), utils.GetDayTime(), menuUid, roleUid).Error
}

// Del 删除
func (u *SysRoleMenu) Del(roleUid string) error {
	return initalize.DB.Exec("delete from sys_role_menu where role_uid = ?", roleUid).Error
}
