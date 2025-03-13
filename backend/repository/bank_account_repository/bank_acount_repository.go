package bank_account_repository

import (
	"backend/models/entity"
)

func (r *BankAccountRepository) ShowAll() ([]entity.BankAcountDetail, error) {
	var BankAccount []entity.BankAcountDetail
	err := r.db.Raw("SELECT bank_accounts.id, bank_accounts.account_number, customers.nama, bank_accounts.bank_amount FROM bank_accounts INNER JOIN customers ON bank_accounts.customer_id = customers.id").Scan(&BankAccount).Error
	return BankAccount, err
}

func (r *BankAccountRepository) ShowById(id uint) ([]entity.BankAcountDetail, error) {
	var BankAccount []entity.BankAcountDetail
	err := r.db.Raw("SELECT bank_accounts.id, bank_accounts.account_number, customers.nama, bank_accounts.bank_amount FROM bank_accounts INNER JOIN customers ON bank_accounts.customer_id = customers.id WHERE bank_accounts.id = ?", id).Scan(&BankAccount).Error
	return BankAccount, err
}

func (r *BankAccountRepository) CreateBankAccount(account entity.BankAccount) ([]entity.BankAcountDetail, error) {
	var BankAccountInsert []entity.BankAcountDetail

	acc_number := account.AccountNumber
	err := r.db.Create(&account).Error
	if err != nil {
		return nil, err
	}
	err2 := r.db.Raw("SELECT bank_accounts.id, bank_accounts.account_number, customers.nama, bank_accounts.bank_amount FROM bank_accounts INNER JOIN customers ON bank_accounts.customer_id = customers.id WHERE bank_accounts.account_number = ?", acc_number).Scan(&BankAccountInsert).Error
	if err2 != nil {
		return nil, err
	}
	return BankAccountInsert, nil
}

func (r *BankAccountRepository) UpdateBankAccount(id uint, account_number *string, bank_amount int) ([]entity.BankAcountDetail, error) {
	var BankAccountUpdate []entity.BankAcountDetail

	err := r.db.Model(&entity.BankAccount{}).Where("id", id).Updates(map[string]any{"account_number": account_number, "bank_amount": bank_amount}).Error
	if err != nil {
		return nil, err
	}
	err2 := r.db.Raw("SELECT bank_accounts.id, bank_accounts.account_number, customers.nama, bank_accounts.bank_amount FROM bank_accounts INNER JOIN customers ON bank_accounts.customer_id = customers.id WHERE bank_accounts.id = ?", id).Scan(&BankAccountUpdate).Error
	if err2 != nil {
		return nil, err
	}
	return BankAccountUpdate, nil
}

func (r *BankAccountRepository) DeleteBankAccount(id uint) error {
	return r.db.Delete(&entity.BankAccount{}, id).Error
}
