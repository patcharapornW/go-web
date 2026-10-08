package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	// สร้าง instance ของ fiber
	app := fiber.New()

	// กำหนด route และ handler
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hi, Go! eiei")
	})

	// start server
	if err := app.Listen(":5599"); err != nil {
		panic(err)
	}

}
