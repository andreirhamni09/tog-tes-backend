package pocket_information_usecase

import (
	"backend/models/entity"
	"backend/repository/pocket_information_repository"
)

type PocketInformationUsecaseInterface interface {
	ShowAll() ([]entity.PocketInformationDetail, error)
	ShowById(bank_account_id uint) ([]entity.PocketInformationDetail, error)
	Create(pocket entity.PocketInformationCreate) (hasil int, serr error)
	Delete(id uint) error
}

type PocketInformationUsecase struct {
	PocketInformationRepository pocket_information_repository.PocketInformationRepositoryInterface
}

func GetPocketInformationUsecase(pocketInformation pocket_information_repository.PocketInformationRepositoryInterface) PocketInformationUsecaseInterface {
	return &PocketInformationUsecase{
		PocketInformationRepository: pocketInformation,
	}
}
