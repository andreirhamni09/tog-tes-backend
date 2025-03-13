package pocket_information_repository

import (
	"backend/models/entity"

	"gorm.io/gorm"
)

type PocketInformationRepositoryInterface interface {
	ShowAll() ([]entity.PocketInformationDetail, error)
	ShowById(bank_account_id uint) ([]entity.PocketInformationDetail, error)
	Create(pocket entity.PocketInformationCreate) (hasil int, serr error)
	Delete(id uint) error
}

type PocketInformationRepository struct {
	db *gorm.DB
}

func GetPocketInformationRepository(DB *gorm.DB) PocketInformationRepositoryInterface {
	return &PocketInformationRepository{
		db: DB,
	}
}
