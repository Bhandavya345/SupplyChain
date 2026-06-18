package models

type Inventory struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	SupplierID  uint    `json:"supplier_id"`
}
