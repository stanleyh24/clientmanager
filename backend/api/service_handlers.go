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

func (a *APIServer) updateService(c *fiber.Ctx) error {
	body := models.Service{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	serviceUpdated, err := a.store.UpdateService(body)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(fiber.Map{"data": serviceUpdated})
}

func (a *APIServer) deleteService(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.store.DeleteService(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(fiber.StatusOK)
}
