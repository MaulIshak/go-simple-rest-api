package repository

import (
	"context"
	"go-simple-rest-api/domain"
)

type customerRepository struct {

}

func NewCustomer() domain.CustomerRepository {
	return &customerRepository{}
}
// Delete implements domain.CustomerRepository.
func (cr *customerRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindAll implements domain.CustomerRepository.
func (cr *customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	panic("unimplemented")
}

// FindById implements domain.CustomerRepository.
func (cr *customerRepository) FindById(ctx context.Context, id string) (domain.Customer, error) {
	panic("unimplemented")
}

// Save implements domain.CustomerRepository.
func (cr *customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	panic("unimplemented")
}

// Update implements domain.CustomerRepository.
func (cr *customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	panic("unimplemented")
}

