package models

type Supplier struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	SupplierName string `json:"supplier_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}
