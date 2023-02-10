package models

import (
	"fiber-nuzn-api/initalize"
	"fiber-nuzn-api/pkg/utils"

	"gorm.io/gorm"
)

// SysUserRole 用户-角色关系表
type SysUserRole struct {
	gorm.Model
	UserUid *string `json:"userUid"`
	RoleUid *string `json:"roleUid"`
}

// NewSysUserRole new一个空的结构体
func NewSysUserRole() *SysUserRole {
	return &SysUserRole{}
}

// GetUserRoleByUerUid 通过 userUid 获取 roleUid
func (u *SysUserRole) GetUserRoleByUerUid(uid string) ([]SysUserRole, error) {
	var s []SysUserRole
	if err := initalize.DB.Raw("select * from sys_user_role where user_uid = ?", uid).Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// Create 创建
func (u *SysUserRole) Create(UserUid, RoleUid string) error {
	return initalize.DB.Exec("insert into sys_user_role (created_at, updated_at, user_uid, role_uid) VALUES (?,?,?,?)",
		utils.GetDayTime(), utils.GetDayTime(), UserUid, RoleUid).Error
}

// Update 修改
func (u *SysUserRole) Update(UserUid, RoleUid string) error {
	return initalize.DB.Exec("update sys_user_role set role_uid = ? where user_uid = ?", RoleUid, UserUid).Error
}

// Del 删除
func (u *SysUserRole) Del(UserUid string) error {
	return initalize.DB.Exec("delete from sys_user_role where user_uid = ?", UserUid).Error
}
