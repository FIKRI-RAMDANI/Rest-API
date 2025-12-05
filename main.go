package main

import (
	"github.com/FIKRI-RAMDANI/Rest-API/internal/api"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/config"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/connection"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/repository"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
