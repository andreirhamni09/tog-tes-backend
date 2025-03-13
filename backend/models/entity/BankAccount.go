package entity

type BankAccount struct {
	Id            uint     `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	CustomerId    uint     ` json:"customer_id"`
	AccountNumber *string  `gorm:"unique;not null" json:"account_number,omitempty"`
	BankAmount    int      `gorm:"not null" json:"bank_amount,omitempty"`
	Customer      Customer `gorm:"references:Id;not null"`
}
type BankAcountDetail struct {
	Id            uint   `json:"id"`
	AccountNumber string `json:"account_number"`
	CustomerNama  string `gorm:"column:nama" json:"customer_nama"`
	BankAmount    int    `json:"bank_amount"`
}

type BankAcountUpdate struct {
	AccountNumber *string `json:"account_number"`
	BankAmount    int     `json:"bank_amount"`
}
