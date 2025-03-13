package pocket_information_repository

import (
	"backend/models/entity"
)

func (r *PocketInformationRepository) ShowAll() ([]entity.PocketInformationDetail, error) {
	var PocketInformationDetail []entity.PocketInformationDetail
	err := r.db.Raw("SELECT p.id as id, c.nama as nama, b.account_number as account_number, p.amount as amount_added, b.bank_amount as total_amount, TO_CHAR(p.tanggal_input, 'YYYY-MM-DD HH24:MM:SS') as tanggal_added FROM pocket_informations as p INNER JOIN bank_accounts as b ON p.bank_account_id = b.id INNER JOIN customers as c ON b.customer_id = c.id").Scan(&PocketInformationDetail).Error
	return PocketInformationDetail, err
}
func (r *PocketInformationRepository) ShowById(bank_account_id uint) ([]entity.PocketInformationDetail, error) {
	var PocketInformationDetail []entity.PocketInformationDetail
	err := r.db.Raw("SELECT p.id as id, c.nama as nama, b.account_number as account_number, p.amount as amount_added, b.bank_amount as total_amount, TO_CHAR(p.tanggal_input, 'YYYY-MM-DD HH24:MM:SS') as tanggal_added FROM pocket_informations as p INNER JOIN bank_accounts as b ON p.bank_account_id = b.id INNER JOIN customers as c ON b.customer_id = c.id WHERE p.bank_account_id = ?", bank_account_id).Scan(&PocketInformationDetail).Error
	return PocketInformationDetail, err
}

func (r *PocketInformationRepository) Create(pocket entity.PocketInformationCreate) (hasil int, serr error) {

	err1 := r.db.Exec("INSERT INTO pocket_informations(bank_account_id, amount, tanggal_input) VALUES (?, ?, ?)", pocket.BankAccountId, pocket.Amount, pocket.TanggalInput).Error
	if err1 != nil {
		return 0, err1
	}
	bank_id := pocket.BankAccountId

	var bank_account []entity.BankAcountDetail
	err2 := r.db.Raw("SELECT * FROM bank_accounts WHERE id = ?", bank_id).Scan(&bank_account).Error

	if err2 != nil {
		serr := err2
		return 0, serr
	}

	amount_all := bank_account[0].BankAmount + pocket.Amount

	err3 := r.db.Exec("UPDATE bank_accounts SET bank_amount = ? WHERE id = ?", amount_all, bank_id).Error

	if err3 != nil {
		return 0, err3
	}

	return amount_all, nil
}

func (r *PocketInformationRepository) Delete(id uint) error {
	return r.db.Delete(&entity.PocketInformation{}, id).Error
}
