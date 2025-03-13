package bank_account_usecase

import "backend/models/entity"

func (r *BankAccountUsecase) ShowAll() ([]entity.BankAcountDetail, error) {
	return r.BankAccountRepository.ShowAll()
}
func (r *BankAccountUsecase) ShowById(id uint) ([]entity.BankAcountDetail, error) {
	return r.BankAccountRepository.ShowById(id)
}
func (r *BankAccountUsecase) CreateBankAccount(bank_account entity.BankAccount) ([]entity.BankAcountDetail, error) {
	return r.BankAccountRepository.CreateBankAccount(bank_account)
}
func (r *BankAccountUsecase) UpdateBankAccount(id uint, account_number *string, bank_amount int) ([]entity.BankAcountDetail, error) {
	return r.BankAccountRepository.UpdateBankAccount(id, account_number, bank_amount)
}
func (r *BankAccountUsecase) DeleteBankAccount(id uint) error {
	return r.BankAccountRepository.DeleteBankAccount(id)
}
