package entity

type Barang struct {
	Post      Post   `gorm:"embedded"`
	Pemakaian string `json:"varian" gorm:"not null"`
}
