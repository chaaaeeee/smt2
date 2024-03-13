package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/users", getUsers)
	app.Post("/createUser", createUser)
	app.Post("/login", loginUser)
	app.Delete("/deleteUser", deleteUser)
	app.Patch("/changePassword", changePassword)
}
