package api

import (
	"context"
	"net/http"
	"time"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/FIKRI-RAMDANI/Rest-API/dto"
	"github.com/FIKRI-RAMDANI/Rest-API/internal/util"
	"github.com/gofiber/fiber/v2"
)

type customerAPi struct {
	costumerService domain.CostumerService
}

func NewCustomer(app *fiber.App, customerService domain.CostumerService) {
	ca := customerAPi{
		costumerService: customerService,
	}
	app.Get("/customers", ca.Index)
	app.Post("/customers", ca.Create)
}

func (ca customerAPi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ca.costumerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ca customerAPi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("validation failed", fails))
	}
	err := ca.costumerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}
