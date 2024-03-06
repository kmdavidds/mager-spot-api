package entity

type Ojek struct {
	Post      Post   `gorm:"embedded"`
	PlatNomor string `json:"platNomor" gorm:"not null"`
}
