package phantom

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

// Config config for middleware
type Config struct {
	Callback func(c fiber.Ctx) (int, error)
}

// Phantom primary object
type Phantom struct {
	Callback func(c fiber.Ctx) (int, error)
}

// New initializes middleware
func New(config *Config) func(c fiber.Ctx) error {
	p := &Phantom{
		Callback: config.Callback,
	}

	return p.middleware
}

// middleware calls callback
func (p *Phantom) middleware(c fiber.Ctx) error {
	if statusCode, err := p.Callback(c); err != nil {
		if statusCode >= 100 || statusCode <= 511 {
			c.Status(statusCode)
		} else {
			c.Status(http.StatusUnauthorized)
		}

		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Next middleware
	return c.Next()
}
