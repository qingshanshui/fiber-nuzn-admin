package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// GET
func CheckQueryParams(c *fiber.Ctx, obj interface{}) error {
	if err := c.QueryParser(obj); err != nil {
		return err
	}
	if err := validateStruct(obj); err != nil {
		return err
	}
	return nil
}

// POST
func CheckPostParams(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return err
	}
	if err := validateStruct(obj); err != nil {
		return err
	}
	return nil
}

// 验证数据
func validateStruct(obj interface{}) error {
	valid := validator.New()
	err := valid.Struct(obj)
	if err != nil {
		return err
	}
	return nil
}
