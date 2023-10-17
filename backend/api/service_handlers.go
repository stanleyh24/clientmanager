package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stanleyh24/clientmanager/models"
)

func (a *APIServer) getAllServices(c *fiber.Ctx) error {
	services, err := a.store.GetAllServices()
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{"data": services})
}

func (a *APIServer) createService(c *fiber.Ctx) error {
	body := models.Service{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	service, err := a.store.CreateService(body)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}

	return c.Status(201).JSON(fiber.Map{"data": service})
}
