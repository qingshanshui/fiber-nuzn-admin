package service

import (
	"fiber-nuzn-api/models"
	v1 "fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (t *UserService) Create(r v1.CreateUserRequest) error {
	m := models.NewSysUser()
	canonical, _ := nanoid.Standard(36)
	uid := canonical()
	m.Uid = uid
	m.Nickname = r.Nickname
	m.Username = r.Username
	m.Password = r.Password
	m.Telephone = r.Telephone
	m.Email = r.Email
	m.Sex = r.Sex
	m.Status = r.Status
	err := m.Create()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "创建用户错误")
	}

	ur := models.NewSysUserRole()
	err = ur.Create(uid, r.Role)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "创建用户角色错误")
	}
	return nil
}

func (t *UserService) Del(r v1.DelUserRequest) error {
	m := models.NewSysUser()
	result, err := m.GetSysUserByUid(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(result) == 0 {
		return fiber.NewError(fiber.StatusServiceUnavailable, "删除用户不存在")
	}
	err = m.Delete(result[0].Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "删除用户错误")
	}
	ur := models.NewSysUserRole()
	err = ur.Del(result[0].Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "删除用户角色错误")
	}
	return nil
}

func (t *UserService) Update(r v1.UpdateUserRequest) error {
	m := models.NewSysUser()
	result, err := m.GetSysUserByUid(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(result) == 0 {
		return fiber.NewError(fiber.StatusServiceUnavailable, "更新用户不存在")
	}
	m.Nickname = r.Nickname
	m.Username = r.Username
	m.Telephone = r.Telephone
	m.Email = r.Email
	m.Sex = r.Sex
	m.Status = r.Status
	err = m.Update(r.Uid)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "更新用户错误")
	}
	ur := models.NewSysUserRole()
	err = ur.Update(result[0].Uid, r.Role)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "更新用户角色错误")
	}
	return nil
}

func (t *UserService) GetList() ([]models.SysUser, error) {
	m := models.NewSysUser()
	result, err := m.GetList()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "用户列表错误")
	}
	return result, nil
}
