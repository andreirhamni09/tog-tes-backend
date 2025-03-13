package entity

import (
	"time"
)

type PocketInformation struct {
	Id            uint        `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	BankAccountId uint        `json:"bank_account_id"`
	Amount        int         `gorm:"not null" json:"amount"`
	TanggalInput  time.Time   `gorm:"not null;type:TIMESTAMP" json:"tanggal_input"`
	BankAccount   BankAccount `gorm:"references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type PocketInformationCreate struct {
	BankAccountId int    `json:"bank_account_id" binding:"required"`
	Amount        int    `gorm:"not null" json:"amount" binding:"required"`
	TanggalInput  string `gorm:"not null" json:"tanggal_input" binding:"required"`
}

type PocketInformationDetail struct {
	Id            uint   `gorm:"column:id" json:"id"`
	Nama          string `gorm:"column:nama" json:"nama"`
	AccountNumber string `gorm:"column:account_number" json:"account_number"`
	AmountAdded   int    `gorm:"column:amount_added" json:"amount_added"`
	TotalAmount   int    `gorm:"column:total_amount" json:"total_amount"`
	TanggalAdded  string `gorm:"column:tanggal_added" json:"tanggal_added"`
}
