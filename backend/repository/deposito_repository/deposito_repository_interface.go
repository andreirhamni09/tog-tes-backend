package deposito_repository

import (
	"backend/models/entity"

	"gorm.io/gorm"
)

type DepositoRepositoryInterface interface {
	ShowAll() ([]entity.DepositoDetail, error)
	ShowById(id uint) ([]entity.DepositoDetail, error)
	CreateDeposito(deposito entity.DepositoCreate) error
}

type DepositoRepository struct {
	db *gorm.DB
}

func GetDepositoRepository(DB *gorm.DB) DepositoRepositoryInterface {
	return &DepositoRepository{
		db: DB,
	}
}
