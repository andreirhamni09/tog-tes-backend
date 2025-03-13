package customer_repository

import (
	"backend/models/entity"

	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	ShowAll() ([]entity.Customer, error)
	ShowById(id uint) ([]entity.Customer, error)
	UpdateCustomer(id uint, username string, nama string, password string) (entity.CustomerUpdate, error)
	CreateCustomer(customer entity.Customer) error
	DeleteCustomer(id uint) error
}

type CustomerRepository struct {
	db *gorm.DB
}

func GetCustomerRepository(DB *gorm.DB) CustomerRepositoryInterface {
	return &CustomerRepository{
		db: DB,
	}
}
