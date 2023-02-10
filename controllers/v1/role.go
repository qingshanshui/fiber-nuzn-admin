package v1

import (
	"fiber-nuzn-api/controllers"
	"fiber-nuzn-api/service"
	"fiber-nuzn-api/validator"
	"fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
)

type RoleController struct {
	controllers.Base
}

func NewRoleController() *RoleController {
	return &RoleController{}
}

func (t *RoleController) Create(c *fiber.Ctx) error {
	// 初始化参数结构体
	r := form.CreateRoleRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &r); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewRoleService().Create(r)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("创建角色成功")) // => ✋ register
}

func (t *RoleController) Del(c *fiber.Ctx) error {
	// 初始化参数结构体
	d := form.DelRoleRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &d); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewRoleService().Del(d)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("删除角色成功")) // => ✋ register
}
func (t *RoleController) Update(c *fiber.Ctx) error {
	// 初始化参数结构体
	u := form.UpdateRoleRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &u); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewRoleService().Update(u)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("编辑角色成功")) // => ✋ register
}
func (t *RoleController) GetList(c *fiber.Ctx) error {
	// 实际业务调用
	result, err := service.NewRoleService().GetList()
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(result)) // => ✋ register
}

// GetCurrentRoleAuthorizationMenu 获取当前角色授权菜单
func (t *RoleController) GetCurrentRoleAuthorizationMenu(c *fiber.Ctx) error {
	u := form.CurrentRoleAuthorizationMenuRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &u); err != nil {
		return err
	}
	// 实际业务调用
	result, err := service.NewRoleService().GetCurrentRoleAuthorizationMenu(u.RoleUid)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(result)) // => ✋ register
}

// SetCurrentRoleAuthorization 给当前角色授权
func (t *RoleController) SetCurrentRoleAuthorization(c *fiber.Ctx) error {
	u := form.SetCurrentRoleAuthorization{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &u); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewRoleService().SetCurrentRoleAuthorization(u)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("授权成功")) // => ✋ register
}
