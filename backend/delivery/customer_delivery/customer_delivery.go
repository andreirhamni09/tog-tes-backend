package customer_delivery

import (
	"backend/models/dto"
	"backend/models/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *CustomerDelivery) ShowAll(c *gin.Context) {
	var response dto.Response
	customer, err := d.CustomerUsecase.ShowAll()
	if err != nil {
		response.Status = 500
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = 200
	response.Messages = "List Customer"
	response.Data = customer
	c.JSON(http.StatusOK, response)
}

func (d *CustomerDelivery) ShowById(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)
	customer, err := d.CustomerUsecase.ShowById(uid)
	if err != nil {
		response.Status = 500
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = 200
	response.Messages = "Detail Customer"
	response.Data = customer
	c.JSON(http.StatusOK, response)
}

func (d *CustomerDelivery) CreateCustomer(c *gin.Context) {
	var customer entity.Customer
	var dataCreated entity.CustomerCreated
	var response dto.Response
	if err := c.ShouldBindJSON(&customer); err != nil {
		response.Status = http.StatusBadRequest
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusOK, response)
		return
	}

	if err := d.CustomerUsecase.CreateCustomer(customer); err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusOK, response)
		return
	}
	dataCreated.Username = customer.Username
	dataCreated.Password = customer.Password
	dataCreated.Nama = customer.Nama

	response.Status = http.StatusOK
	response.Messages = "Customer Berhasil Di Inputkan"
	response.Data = dataCreated
	c.JSON(http.StatusOK, response)
}

func (d *CustomerDelivery) UpdateCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)

	var customer entity.CustomerUpdate
	var response dto.Response
	if err := c.ShouldBindJSON(&customer); err != nil {
		response.Status = http.StatusBadRequest
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	customer, err := d.CustomerUsecase.UpdateCustomer(uid, customer.Username, customer.Nama, customer.Password)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Berhasil Update Data"
	response.Data = customer
	c.JSON(http.StatusOK, response)
}

func (d *CustomerDelivery) DeleteCustomer(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)

	err := d.CustomerUsecase.DeleteCustomer(uid)
	if err != nil {
		response.Status = http.StatusOK
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Data Berhasil Dihapus"
	response.Data = nil
	c.JSON(http.StatusOK, response)
}
