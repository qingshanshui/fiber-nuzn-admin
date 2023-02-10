package v1

import (
	"fiber-nuzn-api/controllers"
	"fiber-nuzn-api/service"
	"fiber-nuzn-api/validator"
	"fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
)

type MenuController struct {
	controllers.Base
}

func NewMenuController() *MenuController {
	return &MenuController{}
}

func (t *MenuController) Create(c *fiber.Ctx) error {
	// 初始化参数结构体
	r := form.CreateMenuRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &r); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewMenuService().Create(r)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("创建菜单成功")) // => ✋ register
}

func (t *MenuController) Del(c *fiber.Ctx) error {
	// 初始化参数结构体
	d := form.DelMenuRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &d); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewMenuService().Del(d)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("删除菜单成功")) // => ✋ register
}
func (t *MenuController) Update(c *fiber.Ctx) error {
	// 初始化参数结构体
	u := form.UpdateMenuRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &u); err != nil {
		return err
	}
	// 实际业务调用
	err := service.NewMenuService().Update(u)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok("编辑菜单成功")) // => ✋ register
}
func (t *MenuController) GetList(c *fiber.Ctx) error {
	// 实际业务调用
	result, err := service.NewMenuService().GetList()
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(result)) // => ✋ register
}
