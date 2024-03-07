package entity

type Barang struct {
	Post      Post   `gorm:"embedded"`
	Pemakaian string `json:"pemakaian" gorm:"not null"`
}
