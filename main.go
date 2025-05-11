package main

import (
	"go-simple-rest-api/dto"
	"go-simple-rest-api/internal/api"
	"go-simple-rest-api/internal/config"
	"go-simple-rest-api/internal/connection"
	"go-simple-rest-api/internal/repository"
	"go-simple-rest-api/internal/service"
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbconnection := connection.GetDatabase(cnf.Database)
	app := fiber.New()

	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{
			Key: []byte(cnf.Jwt.Key),
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error{
			return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("endpoint perlu token, silakan login dulu"))
		},
	})

	customerRepository := repository.NewCustomer(dbconnection)
	userRepository := repository.NewUser(dbconnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(*cnf, userRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}

