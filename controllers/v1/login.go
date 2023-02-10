package v1

import (
	"fiber-nuzn-api/controllers"
	"fiber-nuzn-api/pkg/utils"
	"fiber-nuzn-api/service"
	"fiber-nuzn-api/validator"
	"fiber-nuzn-api/validator/form"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type LoginController struct {
	controllers.Base
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (t *LoginController) Login(c *fiber.Ctx) error {
	// 初始化参数结构体
	r := form.UserRequest{}
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &r); err != nil {
		return err
	}
	// 实际业务调用
	result, err := service.NewLoginService().Login(r)
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(result)) // => ✋ register
}

func (t *LoginController) GetUserInfo(c *fiber.Ctx) error {
	// 解析token 里面的 uid
	token, err := utils.ParseToken(c.GetReqHeaders()["Authorization"])
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	// 拿到 claims结构体（里面存放的token信息）
	claims := token.Claims.(jwt.MapClaims)
	// 实际业务调用
	result, err := service.NewLoginService().GetUserInfo(claims["uid"].(string))
	if err != nil {
		return c.JSON(t.Fail(err))
	}
	return c.JSON(t.Ok(result)) // => ✋ register
}
