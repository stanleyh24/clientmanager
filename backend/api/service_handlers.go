package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/stanleyh24/clientmanager/models"
	"github.com/stanleyh24/clientmanager/utils"
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

	rates := strings.Split(body.Rate, "/")
	rate_limit := fmt.Sprintf("%sk/%sk", rates[0], rates[1])

	q := models.QueryParams{
		Method:   "PUT",
		Resource: "ppp/profile",
		Body:     map[string]string{"name": body.Name, "rate-limit": rate_limit},
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

	body.RouterIdentifier = resp.Body[".id"].(string)

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

	rates := strings.Split(body.Rate, "/")
	rate_limit := fmt.Sprintf("%sk/%sk", rates[0], rates[1])

	q := models.QueryParams{
		Method:   "PATCH",
		Resource: fmt.Sprintf("ppp/profile/%s", body.RouterIdentifier),
		Body:     map[string]string{"name": body.Name, "rate-limit": rate_limit},
	}

	req := utils.Querymaker(q)
	resp, err := utils.SendRequest(req)

	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest)
	}

	if resp.Status != http.StatusOK {
		log.Println(resp)
		return fiber.NewError(fiber.StatusInternalServerError)
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
