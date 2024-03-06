package entity

type Makanan struct {
	Post   Post   `gorm:"embedded"`
	Varian string `json:"varian" gorm:"not null"`
}
