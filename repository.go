package shipment

import (
	"github.com/Bhandavya345/supply-chain-system/internal/database"
	"github.com/Bhandavya345/supply-chain-system/models"
)

func CreateShipment(shipment *models.Shipment) error {
	return database.DB.Create(shipment).Error
}

func GetAllShipments() ([]models.Shipment, error) {

	var shipments []models.Shipment

	err := database.DB.Find(&shipments).Error

	return shipments, err
}

func GetShipmentByID(id uint) (*models.Shipment, error) {

	var shipment models.Shipment

	err := database.DB.First(&shipment, id).Error

	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

func GetShipmentByTrackingNumber(trackingNumber string) (*models.Shipment, error) {

	var shipment models.Shipment

	err := database.DB.
		Where("tracking_number = ?", trackingNumber).
		First(&shipment).Error

	if err != nil {
		return nil, err
	}

	return &shipment, nil
}
func UpdateShipment(id uint, shipment *models.Shipment) error {
	return database.DB.Model(&models.Shipment{}).
		Where("id = ?", id).
		Updates(shipment).Error
}

func DeleteShipment(id uint) error {
	return database.DB.Delete(&models.Shipment{}, id).Error
}
