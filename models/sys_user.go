package models

import (
	"fiber-nuzn-api/initalize"
	"fiber-nuzn-api/pkg/utils"

	"gorm.io/gorm"
)

//
//// 表创建：https://blog.csdn.net/g759780748/article/details/118968658?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1-118968658-blog-113996543.pc_relevant_multi_platform_whitelistv3&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1-118968658-blog-113996543.pc_relevant_multi_platform_whitelistv3&utm_relevant_index=2
//// 参考：https://blog.csdn.net/qq_21852449/article/details/84558843?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1-84558843-blog-121775626.pc_relevant_multi_platform_whitelistv1&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1-84558843-blog-121775626.pc_relevant_multi_platform_whitelistv1&utm_relevant_index=1
//

// SysUser 用户表
type SysUser struct {
	gorm.Model
	Uid       string `json:"uid"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	HeadImg   string `json:"headImg"`
	Sex       string `json:"sex"`
	Age       string `json:"age"`
	Birthday  string `json:"birthday"`
	Status    int    `json:"status"`
	RoleUid   string `json:"role_Uid"`
	//SysUserRole SysUserRole `json:"sysUserRole"`
}

// NewSysUser new一个空的结构体
func NewSysUser() *SysUser {
	return &SysUser{}
}

// Create 创建
func (u *SysUser) Create() error {
	return initalize.DB.Exec("insert into sys_user (created_at, updated_at, uid, username, "+
		"nickname, password, telephone, email, sex, age)"+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", utils.GetDayTime(), utils.GetDayTime(), u.Uid, u.Username,
		u.Nickname, u.Password, u.Telephone, u.Email, u.Sex, u.Age).Error
}

// Delete 删除
func (u *SysUser) Delete(uid string) error {
	return initalize.DB.Exec("delete from sys_user where uid = ?", uid).Error
}

// GetSysUserByUid 通过uid查询 用户 详情
func (u *SysUser) GetSysUserByUid(uid string) ([]SysUser, error) {
	var s []SysUser
	if err := initalize.DB.Raw("select * from sys_user where uid = ?", uid).Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// GetSysUserByUsername 通过 用户名 查询 用户 详情
func (u *SysUser) GetSysUserByUsername(username string) ([]SysUser, error) {
	var s []SysUser
	if err := initalize.DB.Raw("select * from sys_user where username = ?", username).Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// Update 修改
func (u *SysUser) Update(uid string) error {
	return initalize.DB.Exec("update sys_user set updated_at =?,username=?,nickname=?,"+
		"telephone=?,email=?,sex=?,age=? where uid=?", utils.GetDayTime(), u.Username, u.Nickname,
		u.Telephone, u.Email, u.Sex, u.Age, uid).Error
}

// GetList 获取 角色 列表
func (u *SysUser) GetList() ([]SysUser, error) {
	var uArr []SysUser
	if err := initalize.DB.Raw("select * from sys_user left join sys_user_role on sys_user.uid = sys_user_role.user_uid").Find(&uArr).Error; err != nil {
		return nil, err
	}
	return uArr, nil
}
