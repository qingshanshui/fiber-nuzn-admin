package service

import (
	"context"
	"fiber-nuzn-api/initalize"
	"fiber-nuzn-api/models"
	model "fiber-nuzn-api/models"
	"fiber-nuzn-api/pkg/utils"
	v1 "fiber-nuzn-api/validator/form"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (t *LoginService) Login(r v1.UserRequest) (map[string]interface{}, error) {
	m := models.NewSysUser()
	u, err := m.GetSysUserByUsername(r.Username)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(u) == 0 {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "用户不存在")
	}
	// 验证密码 是否正确
	if u[0].Password != r.Password {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "密码错误")
	}
	//  返回的结构体
	rep := v1.UserResponse{
		Uid:      u[0].Uid,
		Nickname: u[0].Nickname,
	}
	// 创建token 传入昵称和uid存到token里面
	token, err := utils.GetToken(u[0].Nickname, u[0].Uid)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "token错误")
	}
	if err := initalize.Rdb.Set(context.Background(), token, true, time.Second*time.Duration(viper.GetInt("Jwt.Expire"))).Err(); err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "redis错误")
	}
	return map[string]interface{}{
		"data":  rep,
		"token": token,
	}, nil
}

func (t *LoginService) GetUserInfo(uid string) ([]model.SysMenu, error) {
	m := model.NewSysUserRole()
	r, err := m.GetUserRoleByUerUid(uid)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "用户数据错误")
	}
	if len(r) == 0 {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "roleUid错误")
	}
	menu := model.NewSysMenu()
	isMenu, err := menu.GetSysMenuByUidIsMenu(*r[0].RoleUid)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, "权限菜单错误")
	}
	return isMenu, nil
}
