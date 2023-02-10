package service

import (
	model "fiber-nuzn-api/models"
	v1 "fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (t *MenuService) Create(r v1.CreateMenuRequest) error {
	m := model.NewSysMenu()
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	m.Uid = uid
	m.Title = r.Title
	m.ParentUid = r.ParentUid
	m.Name = r.Name
	m.Path = r.Path
	m.Component = r.Component
	m.Redirect = r.Redirect
	m.Type = r.Type
	m.Order = r.Order
	err := m.Create()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "创建菜单错误")
	}
	return nil
}

func (t *MenuService) Del(r v1.DelMenuRequest) error {
	m := model.NewSysMenu()
	result, err := m.GetSysMenuByUid(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(result) == 0 {
		return fiber.NewError(fiber.StatusServiceUnavailable, "菜单不存在")
	}
	err = m.Delete(result[0].Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "删除菜单错误")
	}
	return nil
}

func (t *MenuService) Update(r v1.UpdateMenuRequest) error {
	m := model.NewSysMenu()
	result, err := m.GetSysMenuByUid(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(result) == 0 {
		return fiber.NewError(fiber.StatusServiceUnavailable, "菜单不存在")
	}
	m.Title = r.Title
	m.ParentUid = r.ParentUid
	m.Name = r.Name
	m.Path = r.Path
	m.Component = r.Component
	m.Redirect = r.Redirect
	m.Type = r.Type
	m.Order = r.Order
	err = m.Update(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "菜单更新错误")
	}
	return nil
}

func (t *MenuService) GetList() ([]model.SysMenu, error) {
	m := model.NewSysMenu()
	result, err := m.GetList()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "菜单列表错误")
	}
	return result, nil
}
