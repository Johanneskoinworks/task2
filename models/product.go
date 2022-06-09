package models

type Product struct {
	Id        uint    `gorm:"primaryKey" json:"id"`
	Nama      string  `gorm:"not null, type:varchar(191)" json:"nama"`
	Deskripsi string  `gorm:"not null, type:varchar(191)" json:"dekripsi"`
	Harga     uint    `json:"harga"`
	Order     []Order `gorm:"foreignKey:IdProduct; references:Id"`
}
