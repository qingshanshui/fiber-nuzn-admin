package service

import (
	"fiber-nuzn-api/models"
	v1 "fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (t *RoleService) Create(r v1.CreateRoleRequest) error {
	m := models.NewSysRole()
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	m.Uid = uid
	m.Name = r.Name
	m.Remark = r.Remark
	m.Order = r.Order
	m.Status = r.Status
	err := m.Create()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "创建角色错误")
	}
	return nil
}

func (t *RoleService) Del(r v1.DelRoleRequest) error {
	m := models.NewSysRole()
	result, err := m.GetSysRoleByUid(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(result) == 0 {
		return fiber.NewError(fiber.StatusServiceUnavailable, "删除角色不存在")
	}
	err = m.Delete(result[0].Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "删除角色错误")
	}
	return nil
}

func (t *RoleService) Update(r v1.UpdateRoleRequest) error {
	m := models.NewSysRole()
	result, err := m.GetSysRoleByUid(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(result) == 0 {
		return fiber.NewError(fiber.StatusServiceUnavailable, "更新角色不存在")
	}
	m.Name = r.Name
	m.Remark = r.Remark
	m.Order = r.Order
	m.Status = r.Status
	err = m.Update(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "更新角色错误")
	}
	return nil
}

func (t *RoleService) GetList() ([]models.SysRole, error) {
	m := models.NewSysRole()
	result, err := m.GetList()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "角色列表错误")
	}
	return result, nil
}

func (t *RoleService) GetCurrentRoleAuthorizationMenu(roleUid string) ([]models.SysMenu, error) {
	m := models.NewSysMenu()
	result, err := m.GetSysMenuByUidIsMenu(roleUid)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "权限菜单错误")
	}
	return result, nil
}

func (t *RoleService) SetCurrentRoleAuthorization(r v1.SetCurrentRoleAuthorization) error {
	m := models.NewSysRoleMenu()
	err := m.Del(r.RoleUid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "授权菜单错误")
	}
	for _, s := range r.MenuUidList {
		err := m.Create(s, r.RoleUid)
		if err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, "授权菜单错误")
		}
	}

	return nil
}
