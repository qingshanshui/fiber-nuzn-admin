package models

import (
	"fiber-nuzn-api/initalize"
	"fiber-nuzn-api/pkg/utils"

	"gorm.io/gorm"
)

// SysRole 角色表
type SysRole struct {
	gorm.Model
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Order  int    `json:"order"`
	Status int    `json:"status"`
	//SysUserRole SysUserRole `json:"sysUserRole"`
}

// NewSysRole new一个空的结构体
func NewSysRole() *SysRole {
	return &SysRole{}
}

// Create 创建
func (u *SysRole) Create() error {
	return initalize.DB.Exec("insert into sys_role (created_at, updated_at, uid, name, remark, `order`, status)"+
		" VALUES (?,?,?,?,?,?,?)", utils.GetDayTime(), utils.GetDayTime(), u.Uid, u.Name, u.Remark, u.Order, u.Status).Error
}

// Delete 删除
func (u *SysRole) Delete(uid string) error {
	return initalize.DB.Exec("delete from sys_role where uid = ?", uid).Error
}

// GetSysRoleByUid 通过uid查询 用户 详情
func (u *SysRole) GetSysRoleByUid(uid string) ([]SysRole, error) {
	var s []SysRole
	if err := initalize.DB.Raw("select * from sys_role where uid = ?", uid).Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// Update 修改
func (u *SysRole) Update(uid string) error {
	return initalize.DB.Exec("update sys_role set updated_at = ?, name=?,remark=?,`order`=?,status=? where uid = ?",
		utils.GetDayTime(), u.Name, u.Remark, u.Order, u.Status, uid).Error
}

// GetList 获取 角色 列表
func (u *SysRole) GetList() ([]SysRole, error) {
	var uArr []SysRole
	if err := initalize.DB.Raw("select *from sys_role where deleted_at IS NULL order by created_at asc").Find(&uArr).Error; err != nil {
		return nil, err
	}
	return uArr, nil
}

//// GetSysRoleByUidIsUser 通过用户uid查询 查询用户关联的角色 详情
//func (u *SysRole) GetSysRoleByUidIsUser(uid string) *SysRole {
//	if err := u.DB().Joins("SysUserRole").Where("user_uid = ?", uid).First(u).Error; err != nil {
//		return nil
//	}
//	return u
//}
//
