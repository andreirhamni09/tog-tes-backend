package pocket_information_delivery

import (
	"backend/models/dto"
	"backend/models/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (d *PocketInformationDelivery) ShowAll(c *gin.Context) {
	var response dto.Response
	pocket_information, err := d.PocketInformationUsecase.ShowAll()
	if err != nil {
		response.Status = 500
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = 200
	response.Messages = "List Pocket Information"
	response.Data = pocket_information
	c.JSON(http.StatusOK, response)
}

func (d *PocketInformationDelivery) ShowById(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("bank_account_id"))
	uid := uint(id)
	pocket_information, err := d.PocketInformationUsecase.ShowById(uid)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Detail Pocket Information"
	response.Data = pocket_information
	c.JSON(http.StatusOK, response)
}

func (d *PocketInformationDelivery) Create(c *gin.Context) {
	var pocket entity.PocketInformationCreate
	var response dto.Response
	if err := c.ShouldBindJSON(&pocket); err != nil {
		response.Status = http.StatusBadRequest
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err := d.PocketInformationUsecase.Create(pocket)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Messages = "Pocket Information Berhasil Di Inputkan"
	response.Data = nil
	c.JSON(http.StatusCreated, response)
}

func (d *PocketInformationDelivery) Delete(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)
	err := d.PocketInformationUsecase.Delete(uid)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Berhasil Delete Pocket Information"
	response.Data = nil
	c.JSON(http.StatusOK, response)
}
