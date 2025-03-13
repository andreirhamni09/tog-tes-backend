package deposito_repository

import (
	"backend/models/entity"
)

func (r *DepositoRepository) ShowAll() ([]entity.DepositoDetail, error) {
	var DepositoDetail []entity.DepositoDetail
	err := r.db.Raw("SELECT depositos.id as id, depositos.bank_account_id as bank_account_id, bank_accounts.account_number as bank_number, customers.nama as nama, depositos.amount FROM depositos INNER JOIN bank_accounts ON bank_accounts.id = depositos.bank_account_id INNER JOIN customers on customers.id = bank_accounts.customer_id").Scan(&DepositoDetail).Error
	return DepositoDetail, err
}

func (r *DepositoRepository) ShowById(id uint) ([]entity.DepositoDetail, error) {
	var DepositoDetail []entity.DepositoDetail
	err := r.db.Raw("SELECT depositos.id as id, depositos.bank_account_id as bank_account_id, bank_accounts.account_number as bank_number, customers.nama as nama, depositos.amount FROM depositos INNER JOIN bank_accounts ON bank_accounts.id = depositos.bank_account_id INNER JOIN customers on customers.id = bank_accounts.customer_id WHERE depositos.bank_account_id = ?", id).Scan(&DepositoDetail).Error
	return DepositoDetail, err
}

func (r *DepositoRepository) CreateDeposito(deposito entity.DepositoCreate) error {
	err1 := r.db.Exec("INSERT INTO depositos(bank_account_id, amount) VALUES (?,?)", deposito.BankAccountId, deposito.Amount).Error
	if err1 != nil {
		return err1
	}
	return nil
}
