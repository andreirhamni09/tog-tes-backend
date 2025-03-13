package customer_usecase

import "backend/models/entity"

func (cus *CustomerUsecase) ShowAll() ([]entity.Customer, error) {
	return cus.CustomerRepository.ShowAll()
}
func (cus *CustomerUsecase) ShowById(id uint) ([]entity.Customer, error) {
	return cus.CustomerRepository.ShowById(id)
}
func (cus *CustomerUsecase) CreateCustomer(customer entity.Customer) error {
	return cus.CustomerRepository.CreateCustomer(customer)
}
func (cus *CustomerUsecase) UpdateCustomer(id uint, username string, nama string, password string) (entity.CustomerUpdate, error) {
	return cus.CustomerRepository.UpdateCustomer(id, username, nama, password)
}
func (cus *CustomerUsecase) DeleteCustomer(id uint) error {
	return cus.CustomerRepository.DeleteCustomer(id)
}
