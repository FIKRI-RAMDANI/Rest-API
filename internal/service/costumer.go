package service

import (
	"context"
	"database/sql"
	"errors"
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

func (c customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := c.costumerRepository.FindById(ctx, id)
	if err != nil {
		return dto.CustomerData{}, err
	}
	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("Data customer tidak ditemukan")
	}
	return dto.CustomerData{
		ID:   persisted.ID,
		Code: persisted.Code,
		Name: persisted.Name,
	}, nil
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

func (c customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.costumerRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("Data customer tidak ditemukan")
	}
	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return c.costumerRepository.Update(ctx, &persisted)
}

func (c customerService) Delete(ctx context.Context, id string) error {
	exits, err := c.costumerRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exits.ID == "" {
		return errors.New("Data customer tidak ditemukan")
	}
	return c.costumerRepository.Delete(ctx, id)
}
