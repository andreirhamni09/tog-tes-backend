package entity

type Customer struct {
	Id       uint   `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

type CustomerUpdate struct {
	Username string `gorm:"unique" json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}
type CustomerCreated struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}
