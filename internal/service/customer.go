package service

import (
	"context"
	"database/sql"
	"go-simple-rest-api/domain"
	"go-simple-rest-api/dto"
	"time"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

// Create implements domain.CustomerService.

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

// Index implements domain.CustomerService.
func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customersData []dto.CustomerData

	for _, val := range customers {
		customersData = append(customersData, dto.CustomerData{
			ID:   val.ID,
			Code: val.Code,
			Name: val.Name,
		})
	}
	return customersData, nil
}

func (c *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID: uuid.NewString(),
		Code: req.Code,
		Name: req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return c.customerRepository.Save(ctx, &customer)
}