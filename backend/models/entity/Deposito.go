package entity

import "time"

type Deposito struct {
	Id            uint        `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	BankAccountId uint        `json:"bank_account_id"`
	Amount        int         `gorm:"not null" json:"amount"`
	BankAccount   BankAccount `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type DepositoCreate struct {
	BankAccountId uint `json:"bank_account_id" binding:"required"`
	Amount        int  `json:"amount"`
}

type DepositoDetail struct {
	// id, bank_account_id, bank_number, nama, amount
	Id         uint   `gorm:"column:id" json:"id"`
	Nama       string `gorm:"column:nama" json:"nama"`
	BankNumber string `gorm:"column:bank_number" json:"bank_number"`
	Amount     int    `gorm:"column:amount" json:"amount"`
}

type DepositoHistory struct {
	Id           uint      `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	DepositoId   uint      `json:"deposito_id"`
	Amount       int       `gorm:"not null" json:"amount"`
	TanggalInput time.Time `gorm:"not null;type:TIMESTAMP" json:"tanggal_input"`
	Deposito     Deposito  `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type DepositoHistoryDetail struct {
	Id            uint   `gorm:"column:id" json:"id"`
	Nama          string `gorm:"column:nama" json:"nama"`
	AccountNumber string `gorm:"column:account_number" json:"account_number"`
	AmountAdded   int    `gorm:"column:amount_added" json:"amount_added"`
	TanggalAdded  string `gorm:"column:tanggal_added" json:"tanggal_added"`
}
