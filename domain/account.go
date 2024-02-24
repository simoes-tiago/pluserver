package domain

type Account struct {
	ID      int64 `gorm:"primaryKey"`
	UserID  int64
	Balance float32
}
