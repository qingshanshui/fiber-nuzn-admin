package v1

import (
	"fiber-nuzn-api/controllers"
	"fiber-nuzn-api/service"
	"fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
)

type TreeController struct {
	controllers.Base
}

func NewTreeController() *TreeController {
	return &TreeController{}
}

// GetList 获取角色
func (t *TreeController) GetList(ctx *fiber.Ctx) error {
	// 初始化参数结构体
	var TreesStruct form.TreesStruct
	if err := ctx.BodyParser(&TreesStruct); err != nil {
		return ctx.JSON(t.Fail(err))
	}
	// 实际业务调用
	result, err := service.NewTreeService().GetList(TreesStruct)
	// 根据业务返回值判断业务成功 OR 失败
	if err != nil {
		return ctx.JSON(t.Fail(err))
	}
	return ctx.JSON(t.Ok(result))
}
