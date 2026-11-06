package main

import (
	"codeberg.org/invirtuate/phantom"
	"codeberg.org/invirtuate/phantom/internal/ipban"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Use(phantom.New(&phantom.Config{
		Callback: func(c fiber.Ctx) (int, error) {
			// check ip banned
			if err := ipban.Authorized("1.1.1.1"); err != nil {
				return fiber.StatusForbidden, err
			}

			return fiber.StatusOK, nil
		},
	}))

	app.Listen(":3000")
}
