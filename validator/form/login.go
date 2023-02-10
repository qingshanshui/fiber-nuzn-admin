package form

import "gorm.io/gorm"

// UserRequest 登录
type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse 登录相应
type UserResponse struct {
	Uid      string `json:"uid"`      // 用户唯一的uid
	Nickname string `json:"nickname"` // 昵称
	Roles    string `json:"roles"`    // 权限的uid
}

// UserInfoRequest 获取用户信息
type UserInfoRequest struct {
	Uid string `json:"uid" validate:"required"`
}

// UserMenuRequest 用户信息返回数据
type UserMenuRequest struct {
	gorm.Model
	Title     string `json:"title" validate:"required"`     // 名称
	Path      string `json:"path" validate:"required"`      // 路由路径
	Redirect  string `json:"redirect"`                      // 重定向
	Name      string `json:"name" validate:"required"`      // 组件名称
	Component string `json:"component" validate:"required"` // 组件
	Order     int    `json:"order"`                         // 排序
	Type      string `json:"type" validate:"required"`      // 类型
	ParentUid string `json:"parentUid" validate:"required"` // 父级uid
}
