package domain

type User struct {
	ID       int64    `json:",omitempty" gorm:"primaryKey;autoincrement"`
	Username string   `gorm:"unique;not null;default:null"`
	Phone    string   `json:",omitempty"`
	Name     string   `json:",omitempty"`
	Country  string   `json:",omitempty"`
	Age      int      `json:",omitempty"`
	Account  *Account `json:",omitempty"`
}
