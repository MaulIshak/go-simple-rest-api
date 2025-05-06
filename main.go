package main

import (
	"go-simple-rest-api/internal/config"
	"go-simple-rest-api/internal/connection"
	"go-simple-rest-api/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbconnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	app.Get("/developers", developers)

	customerRepository := repository.NewCustomer(dbconnection)
	_ = app.Listen(cnf.Server.Host + ":"+cnf.Server.Port)
}

func developers(ctx *fiber.Ctx) error{
	return ctx.Status(200).JSON("data")
}