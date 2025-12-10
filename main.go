package main

import (
	"net/http"

	"github.com/FIKRI-RAMDANI/Rest-API/dto"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/api"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/config"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/connection"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/repository"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/service"
	"github.com/gofiber/fiber/v2"
	jwtMid "github.com/gofiber/jwt/v3"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey: []byte(cnf.Jwt.Key),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).JSON(dto.CreateResponseError("iternal server error"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)
	bookRepository := repository.NewBook(dbConnection)
	bookStockRepository := repository.NewStock(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)
	bookService := service.NewBook(bookRepository, bookStockRepository)
	bookStockService := service.NewBookStock(bookRepository, bookStockRepository)

	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)
	api.NewBook(app, bookService, jwtMidd)
	api.NewBookStock(app, bookStockService, jwtMidd)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
