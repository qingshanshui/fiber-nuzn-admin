package form

import "gorm.io/gorm"

// CreateMenuRequest 创建菜单request
type CreateMenuRequest struct {
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

// DelMenuRequest 删除菜单request
type DelMenuRequest struct {
	Uid string `json:"uid" validate:"required"` // uid
}

// UpdateMenuRequest 编辑菜单request
type UpdateMenuRequest struct {
	gorm.Model
	Title     string `json:"title" validate:"required"`     // 名称
	Path      string `json:"path" validate:"required"`      // 路由路径
	Redirect  string `json:"redirect"`                      // 重定向
	Name      string `json:"name" validate:"required"`      // 组件名称
	Component string `json:"component" validate:"required"` // 组件
	Order     int    `json:"order"`                         // 排序
	Type      string `json:"type" validate:"required"`      // 类型
	Uid       string `json:"uid" validate:"required"`       // uid
	ParentUid string `json:"parentUid" validate:"required"` // 父级uid
}
