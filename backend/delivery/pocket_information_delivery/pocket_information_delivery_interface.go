package pocket_information_delivery

import (
	"backend/usecase/pocket_information_usecase"

	"github.com/gin-gonic/gin"
)

type PocketInformationDeliveryInterface interface {
	ShowAll(c *gin.Context)
	ShowById(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type PocketInformationDelivery struct {
	PocketInformationUsecase pocket_information_usecase.PocketInformationUsecaseInterface
}

func GetPocketInformationDelivery(pocketInformationUsecase pocket_information_usecase.PocketInformationUsecaseInterface) PocketInformationDeliveryInterface {
	return &PocketInformationDelivery{
		PocketInformationUsecase: pocketInformationUsecase,
	}
}
