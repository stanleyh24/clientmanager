package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/stanleyh24/clientmanager/models"
	"github.com/stanleyh24/clientmanager/utils"
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

	var userData map[string]interface{}
	err := json.Unmarshal(body.User_Data, &userData)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	q := models.QueryParams{
		Method:   "PUT",
		Resource: "ppp/secret",
		Body:     map[string]string{"name": userData["username"].(string), "password": userData["password"].(string), "profile": userData["profile"].(string)},
	}

	req := utils.Querymaker(q)
	resp, err := utils.SendRequest(req)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}

	if resp.Status != http.StatusCreated {
		log.Println(resp)
		return fiber.NewError(fiber.StatusInternalServerError)
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
