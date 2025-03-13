package pocket_information_usecase

import "backend/models/entity"

func (poc *PocketInformationUsecase) ShowAll() ([]entity.PocketInformationDetail, error) {
	return poc.PocketInformationRepository.ShowAll()
}

func (poc *PocketInformationUsecase) ShowById(bank_account_id uint) ([]entity.PocketInformationDetail, error) {
	return poc.PocketInformationRepository.ShowById(bank_account_id)
}
func (poc *PocketInformationUsecase) Create(pocket entity.PocketInformationCreate) (hasil int, serr error) {
	return poc.PocketInformationRepository.Create(pocket)
}

func (poc *PocketInformationUsecase) Delete(id uint) error {
	return poc.PocketInformationRepository.Delete(id)
}
