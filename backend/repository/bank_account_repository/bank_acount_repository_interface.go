package bank_account_repository

import (
	"backend/models/entity"

	"gorm.io/gorm"
)

type BankAccountRepositoryInterface interface {
	ShowAll() ([]entity.BankAcountDetail, error)
	ShowById(id uint) ([]entity.BankAcountDetail, error)
	CreateBankAccount(account entity.BankAccount) ([]entity.BankAcountDetail, error)
	UpdateBankAccount(id uint, account_number *string, bank_amount int) ([]entity.BankAcountDetail, error)
	DeleteBankAccount(id uint) error
}

type BankAccountRepository struct {
	db *gorm.DB
}

func GetBankAccountRepository(DB *gorm.DB) BankAccountRepositoryInterface {
	return &BankAccountRepository{
		db: DB,
	}
}
