package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/FIKRI-RAMDANI/Rest-API/dto"
	"github.com/google/uuid"
)

type customerService struct {
	costumerRepository domain.CustomerRepository
}

func NewCustomer(costumerRepository domain.CustomerRepository) domain.CostumerService {
	return &customerService{
		costumerRepository: costumerRepository,
	}
}

func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.costumerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}
	return customerData, nil
}

func (c customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return c.costumerRepository.Save(ctx, &customer)
}
