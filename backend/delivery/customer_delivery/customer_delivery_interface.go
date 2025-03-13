package customer_delivery

import (
	"backend/usecase/customer_usecase"

	"github.com/gin-gonic/gin"
)

type CustomerDeliveryInterface interface {
	ShowAll(c *gin.Context)
	ShowById(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type CustomerDelivery struct {
	CustomerUsecase customer_usecase.CustomerUsecaseInterface
}

func GetCustomerDelivery(customerUsecase customer_usecase.CustomerUsecaseInterface) CustomerDeliveryInterface {
	return &CustomerDelivery{
		CustomerUsecase: customerUsecase,
	}
}
