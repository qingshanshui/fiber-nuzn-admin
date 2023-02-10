package v1

import (
	"fiber-nuzn-api/controllers"
	"fiber-nuzn-api/service"
	"fiber-nuzn-api/validator"
	"fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	controllers.Base
}

func NewUserController() *UserController {
	return &UserController{}
}

func (t *UserController) Create(c *fiber.Ctx) error {
	// 初始化参数结构体
	r := form.CreateUserRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &r); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewUserService().Create(r)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("创建用户成功")) // => ✋ register
}

func (t *UserController) Del(c *fiber.Ctx) error {
	// 初始化参数结构体
	d := form.DelUserRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &d); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewUserService().Del(d)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("删除用户成功")) // => ✋ register
}
func (t *UserController) Update(c *fiber.Ctx) error {
	// 初始化参数结构体
	u := form.UpdateUserRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &u); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewUserService().Update(u)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("编辑用户成功")) // => ✋ register
}
func (t *UserController) GetList(c *fiber.Ctx) error {
	// 实际业务调用
	result, err := service.NewUserService().GetList()
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(result)) // => ✋ register
}
