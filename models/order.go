package models

type Order struct {
	Id           uint `gorm:"primaryKey json:"id""`
	IdProduct    uint `json:"id_product"`
	JumlahProduk uint `json:"jumlah_produk"`
	// product      Product `gorm:"foreignKey:IdProduct"`
}
