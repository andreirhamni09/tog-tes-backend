package bank_account_delivery

import (
	"backend/models/dto"
	"backend/models/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (b *BankAccountDelivery) ShowAll(c *gin.Context) {
	var response dto.Response
	bank_account, err := b.BankAccountUsecase.ShowAll()
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Data Bank Account"
	response.Data = bank_account
	c.JSON(http.StatusOK, response)
}

func (b *BankAccountDelivery) ShowById(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)
	bank_account, err := b.BankAccountUsecase.ShowById(uid)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Data Bank Account"
	response.Data = bank_account
	c.JSON(http.StatusOK, response)
}

func (b *BankAccountDelivery) CreateBankAccount(c *gin.Context) {
	var bank_account entity.BankAccount
	var dataCreated []entity.BankAcountDetail
	var response dto.Response
	if err := c.ShouldBindJSON(&bank_account); err != nil {
		response.Status = http.StatusBadRequest
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	insert, err := b.BankAccountUsecase.CreateBankAccount(bank_account)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	dataCreated = insert
	response.Status = 200
	response.Messages = "Customer Berhasil Di Inputkan"
	response.Data = dataCreated
	c.JSON(http.StatusCreated, response)
}

func (b *BankAccountDelivery) UpdateBankAccount(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)

	var update_data entity.BankAcountUpdate

	if err := c.ShouldBindJSON(&update_data); err != nil {
		response.Status = http.StatusBadRequest
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusBadRequest, response)
		return
	}
	bank_account, err := b.BankAccountUsecase.UpdateBankAccount(uid, update_data.AccountNumber, update_data.BankAmount)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Data Berhasil Diupdate"
	response.Data = bank_account
	c.JSON(http.StatusOK, response)
}

func (b *BankAccountDelivery) DeleteBankAccount(c *gin.Context) {
	var response dto.Response
	id, _ := strconv.Atoi(c.Param("id"))
	uid := uint(id)
	err := b.BankAccountUsecase.DeleteBankAccount(uid)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Messages = err.Error()
		response.Data = nil
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Messages = "Data Bank Account Berhasil Dihapus"
	response.Data = nil
	c.JSON(http.StatusOK, response)
}
