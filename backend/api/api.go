package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stanleyh24/clientmanager/storage"
)

type APIServer struct {
	store storage.PostgresStore
}

func NewAPIServer(store storage.PostgresStore) *APIServer {
	return &APIServer{
		store: store,
	}
}

func (a *APIServer) Run() {
	app := fiber.New()

	v1 := app.Group("api/public/v1")

	router := v1.Group("/routers")
	router.Get("/", a.getAllRouter)
	router.Post("/", a.createRouter)
	/*	router.Get("/:id", a.getRouterByID)

		router.Put("/:id", a.updateRouter)
		router.Delete("/:id", a.deleteRouter) */

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":" + os.Getenv("SERVER_PORT")))
}
