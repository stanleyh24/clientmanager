package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stanleyh24/clientmanager/models"
)

func (a *APIServer) getAllRouter(c *fiber.Ctx) error {
	routers, err := a.store.GetAllRouter()
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(fiber.Map{"data": routers})
}

func (a *APIServer) createRouter(c *fiber.Ctx) error {
	body := models.Router{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	router, err := a.store.CreateRouter(body)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}

	return c.Status(201).JSON(fiber.Map{"data": router})
}

func (a *APIServer) updateRouter(c *fiber.Ctx) error {
	body := models.Router{}

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	routerUpdated, err := a.store.UpdateRouter(body)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(fiber.Map{"data": routerUpdated})
}

func (a *APIServer) deleteRouter(c *fiber.Ctx) error {
	id := c.Params("id")
	err := a.store.DeleteRouter(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(fiber.StatusOK)
}
