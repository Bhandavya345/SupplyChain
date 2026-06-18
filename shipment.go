package models

import "time"

type Shipment struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	TrackingNumber string    `gorm:"unique" json:"tracking_number"`
	InventoryID    uint      `json:"inventory_id"`
	Source         string    `json:"source"`
	Destination    string    `json:"destination"`
	Status         string    `json:"status"`
	ShipmentDate   time.Time `json:"shipment_date"`
	DeliveryDate   time.Time `json:"delivery_date"`
}
