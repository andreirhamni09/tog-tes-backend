package bank_account_usecase

import (
	"backend/models/entity"
	"backend/repository/bank_account_repository"
)

type BankAccountUsecaseInterface interface {
	ShowAll() ([]entity.BankAcountDetail, error)
	ShowById(id uint) ([]entity.BankAcountDetail, error)
	CreateBankAccount(bank_account entity.BankAccount) ([]entity.BankAcountDetail, error)
	UpdateBankAccount(id uint, account_number *string, bank_amount int) ([]entity.BankAcountDetail, error)
	DeleteBankAccount(id uint) error
}

type BankAccountUsecase struct {
	BankAccountRepository bank_account_repository.BankAccountRepositoryInterface
}

func GetBankAccountUsecase(bankAccountRepository bank_account_repository.BankAccountRepositoryInterface) BankAccountUsecaseInterface {
	return &BankAccountUsecase{
		BankAccountRepository: bankAccountRepository,
	}
}
