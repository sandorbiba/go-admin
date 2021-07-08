package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sandorbiba/go-admin/models"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Almafa",
	}

	user.LastName = "Nagy"

	return c.JSON(user)
}