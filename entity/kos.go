package entity

type Kos struct {
	Post       Post   `gorm:"embedded"`
	Pembayaran string `json:"pembayaran" gorm:"not null"`
}
