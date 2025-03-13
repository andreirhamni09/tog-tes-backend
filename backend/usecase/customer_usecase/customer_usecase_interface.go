package customer_usecase

import (
	"backend/models/entity"
	"backend/repository/customer_repository"
)

type CustomerUsecaseInterface interface {
	ShowAll() ([]entity.Customer, error)
	ShowById(id uint) ([]entity.Customer, error)
	CreateCustomer(customer entity.Customer) error
	UpdateCustomer(id uint, username string, nama string, password string) (entity.CustomerUpdate, error)
	DeleteCustomer(id uint) error
}

type CustomerUsecase struct {
	CustomerRepository customer_repository.CustomerRepositoryInterface
}

func GetCustomerUsecase(customerRepository customer_repository.CustomerRepositoryInterface) CustomerUsecaseInterface {
	return &CustomerUsecase{
		CustomerRepository: customerRepository,
	}
}
