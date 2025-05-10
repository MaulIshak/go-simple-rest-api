package service

import (
	"context"
	"go-simple-rest-api/domain"
	"go-simple-rest-api/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}


func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

// Index implements domain.CustomerService.
func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil{
		return nil, err
	}

	var customersData	[]dto.CustomerData
	
	for _, val := range customers{
		customersData = append(customersData, dto.CustomerData{
			ID: val.ID,
			Code: val.Code,
			Name: val.Name,
		})
	}	
	return customersData, nil
}