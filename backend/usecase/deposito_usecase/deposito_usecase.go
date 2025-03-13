package deposito_usecase

import "backend/models/entity"

func (d *DepositoUsecase) ShowAll() ([]entity.DepositoDetail, error) {
	return d.DepositoRepository.ShowAll()
}
func (d *DepositoUsecase) ShowById(id uint) ([]entity.DepositoDetail, error) {
	return d.DepositoRepository.ShowById(id)
}
func (d *DepositoUsecase) CreateDeposito(deposito entity.DepositoCreate) error {
	return d.DepositoRepository.CreateDeposito(deposito)
}
