package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stanleyh24/clientmanager/models"
)

func (a *APIServer) getAllClients(c *fiber.Ctx) error {
	routers, err := a.store.GetAllClients()
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{"data": routers})
}

func (a *APIServer) createClient(c *fiber.Ctx) error {
	body := models.Client{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	client, err := a.store.CreateClient(body)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}

	return c.Status(201).JSON(fiber.Map{"data": client})
}

func (a *APIServer) updateClient(c *fiber.Ctx) error {
	body := models.Client{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	clientUpdated, err := a.store.UpdateClient(body)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(fiber.Map{"data": clientUpdated})
}

func (a *APIServer) deleteClient(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.store.DeleteClient(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(fiber.StatusOK)
}
