package customer_repository

import "backend/models/entity"

func (r *CustomerRepository) ShowAll() ([]entity.Customer, error) {
	var customer []entity.Customer
	err := r.db.Find(&customer).Error
	return customer, err
}

func (r *CustomerRepository) ShowById(id uint) ([]entity.Customer, error) {
	var customer []entity.Customer
	err := r.db.First(&customer, id).Error
	return customer, err
}

func (r *CustomerRepository) CreateCustomer(customer entity.Customer) error {
	return r.db.Create(&customer).Error
}

func (r *CustomerRepository) UpdateCustomer(id uint, username string, nama string, password string) (entity.CustomerUpdate, error) {
	var customer entity.CustomerUpdate
	customer.Username = username
	customer.Nama = nama
	customer.Password = password
	err := r.db.Model(&entity.Customer{}).Where("id", id).Updates(map[string]interface{}{"username": username, "nama": nama, "password": password}).Error
	return customer, err
}

func (r *CustomerRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&entity.Customer{}, id).Error
}
