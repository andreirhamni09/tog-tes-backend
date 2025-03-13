package deposito_usecase

import (
	"backend/models/entity"
	"backend/repository/deposito_repository"
)

type DepositoUsecaseInterface interface {
	ShowAll() ([]entity.DepositoDetail, error)
	ShowById(id uint) ([]entity.DepositoDetail, error)
	CreateDeposito(deposito entity.DepositoCreate) error
}

type DepositoUsecase struct {
	DepositoRepository deposito_repository.DepositoRepositoryInterface
}

func GetDepositoUsecase(Deposito deposito_repository.DepositoRepositoryInterface) DepositoUsecaseInterface {
	return &DepositoUsecase{
		DepositoRepository: Deposito,
	}
}
