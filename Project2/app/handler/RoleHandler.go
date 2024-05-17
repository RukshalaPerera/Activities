package handler

import (
	"Project2/app/db"
	"Project2/app/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllRoles(c *fiber.Ctx) error {
	var Roles []model.Role
	db.DB.Find(&Roles)
	return c.JSON(Roles)
}

func CreateRole(c *fiber.Ctx) error {
	var Role model.Role
	if err := c.BodyParser(&Role); err != nil {
		return err
	}
	db.DB.Create(&Role)
	return c.JSON(Role)
}
