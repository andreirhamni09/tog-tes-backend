package deposito_delivery

import (
	"backend/usecase/deposito_usecase"

	"github.com/gin-gonic/gin"
)

type DepositoDeliveryInterface interface {
	ShowAll(c *gin.Context)
	ShowById(c *gin.Context)
	CreateDeposito(c *gin.Context)
}

type DepositoDelivery struct {
	DepositoUsecase deposito_usecase.DepositoUsecaseInterface
}

func GetDepositoDelivery(depositoUsecase deposito_usecase.DepositoUsecaseInterface) DepositoDeliveryInterface {
	return &DepositoDelivery{
		DepositoUsecase: depositoUsecase,
	}
}
