package main

import (
	"go-simple-rest-api/internal/api"
	"go-simple-rest-api/internal/config"
	"go-simple-rest-api/internal/connection"
	"go-simple-rest-api/internal/repository"
	"go-simple-rest-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbconnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	customerRepository := repository.NewCustomer(dbconnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

