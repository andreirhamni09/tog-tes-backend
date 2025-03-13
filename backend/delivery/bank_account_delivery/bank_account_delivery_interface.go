package bank_account_delivery

import (
	"backend/usecase/bank_account_usecase"

	"github.com/gin-gonic/gin"
)

type BankAccountDeliveryInterface interface {
	ShowAll(c *gin.Context)
	ShowById(c *gin.Context)
	CreateBankAccount(c *gin.Context)
	UpdateBankAccount(c *gin.Context)
	DeleteBankAccount(c *gin.Context)
}

type BankAccountDelivery struct {
	BankAccountUsecase bank_account_usecase.BankAccountUsecaseInterface
}

func GetBankAccountDelivery(bankAccountUsecase bank_account_usecase.BankAccountUsecaseInterface) BankAccountDeliveryInterface {
	return &BankAccountDelivery{
		BankAccountUsecase: bankAccountUsecase,
	}
}
