package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wortiz1027/golang/fiber_crud_api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hello", handlers.Hello)
}
