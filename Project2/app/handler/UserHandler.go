package handler

import (
	"Project2/app/db"
	"Project2/app/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []model.User
	if err := db.DB.Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var newUser model.User
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}
	// Check ID exists
	var existingRole model.Role
	result := db.DB.First(&existingRole, newUser.RoleID)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Role not found")
	}
	db.DB.Create(&newUser)
	return c.JSON(newUser)
}

func UpdateUser(c *fiber.Ctx) error {
	UserID := c.Params("id")
	var user model.User
	if err := db.DB.First(&user, UserID).Error; err != nil {
		return err
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	UserID := c.Params("id")
	var user model.User
	if err := db.DB.First(&user, UserID).Error; err != nil {
		return err
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		return err
	}
	return c.SendString("User Deleted")
}
