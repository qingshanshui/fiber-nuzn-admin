package models

import (
	"fiber-nuzn-api/initalize"
	"fiber-nuzn-api/pkg/utils"

	"gorm.io/gorm"
)

// SysMenu 菜单表
type SysMenu struct {
	gorm.Model
	Title     string `json:"title" `
	Uid       string `json:"uid"`
	ParentUid string `json:"parentUid"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Icon      string `json:"icon"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Order     int    `json:"order"`
}

// NewSysMenu new一个空的结构体
func NewSysMenu() *SysMenu {
	return &SysMenu{}
}

// Create 创建
func (u *SysMenu) Create() error {
	return initalize.DB.Exec("INSERT INTO sys_menu (created_at, updated_at, deleted_at, title, uid, "+
		"name, path, component, redirect, "+
		"icon, type,status, `order`)"+
		"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)", utils.GetDayTime(), utils.GetDayTime(), nil, u.Title, u.Uid, u.Name,
		u.Path, u.Component, u.Redirect, u.Icon, u.Type, u.Status, u.Order).Error
}

// Delete 删除
func (u *SysMenu) Delete(uid string) error {
	return initalize.DB.Exec("delete from sys_menu where uid = ?", uid).Error
}

// GetSysMenuByUid 通过uid查询 菜单 详情
func (u *SysMenu) GetSysMenuByUid(uid string) ([]SysMenu, error) {
	var s []SysMenu
	if err := initalize.DB.Raw("select * from sys_menu where uid = ?", uid).Find(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// Update 修改
func (u *SysMenu) Update(uid string) error {
	return initalize.DB.Exec("update sys_menu set updated_at = ? ,title=?,name =? ,path =?,"+
		"component=?,redirect=?,icon=?,type=? where uid = ?", utils.GetDayTime(), u.Title, u.Name, u.Path,
		u.Component, u.Redirect, u.Icon, u.Type, uid).Error
}

// GetList 获取 菜单 列表
func (u *SysMenu) GetList() ([]SysMenu, error) {
	var uArr []SysMenu
	if err := initalize.DB.Raw("select *from sys_menu order by created_at asc").Find(&uArr).Error; err != nil {
		return nil, err
	}
	return uArr, nil
}

// GetSysMenuByUidIsMenu 通过角色uid查询 关联菜单详情
func (u *SysMenu) GetSysMenuByUidIsMenu(RoleUid string) ([]SysMenu, error) {
	var m []SysMenu
	if err := initalize.DB.Raw("select sys_menu.path,sys_menu.name,sys_menu.title,sys_menu.component,"+
		"sys_menu.redirect,sys_menu.`order`,sys_menu.type,sys_menu.uid,sys_menu.parent_uid from sys_menu left "+
		"join sys_role_menu on sys_menu.uid =sys_role_menu.menu_uid where role_uid = ? and sys_menu.deleted_at is null and sys_role_menu.deleted_at is null", RoleUid).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
