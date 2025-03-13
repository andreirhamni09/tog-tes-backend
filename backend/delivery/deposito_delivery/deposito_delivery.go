package deposito_delivery

import (
	"backend/models/dto"
	"backend/models/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *DepositoDelivery) ShowAll(c *gin.Context) {
	var response dto.Response
	deposito, err := d.DepositoUsecase.ShowAll()
	if err != nil {
		response.Status = 500
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = 200
	response.Messages = "List Deposito"
	response.Data = deposito
	c.JSON(http.StatusOK, response)
}

func (d *DepositoDelivery) ShowById(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)
	deposito, err := d.DepositoUsecase.ShowById(uid)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusOK, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Detail Deposito"
	response.Data = deposito
	c.JSON(http.StatusOK, response)
}

func (d *DepositoDelivery) CreateDeposito(c *gin.Context) {

	var deposito entity.DepositoCreate
	var response dto.Response
	if err := c.ShouldBindJSON(&deposito); err != nil {
		response.Status = http.StatusBadRequest
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err := d.DepositoUsecase.CreateDeposito(deposito)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Messages = "Account Deposito Berhasil Dibuat"
	response.Data = nil
	c.JSON(http.StatusCreated, response)
}
