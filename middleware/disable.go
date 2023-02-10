package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Disable(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"code": 0,
		"data": "主体数据，不能变动",
		"msg":  "操作失败",
	})
	//return ctx.Next()
}
