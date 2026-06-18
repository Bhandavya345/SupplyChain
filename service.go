package shipment

import "github.com/Bhandavya345/supply-chain-system/models"

func AddShipment(shipment *models.Shipment) error {
	return CreateShipment(shipment)
}

func FetchAllShipments() ([]models.Shipment, error) {
	return GetAllShipments()
}

func FetchShipmentByID(id uint) (*models.Shipment, error) {
	return GetShipmentByID(id)
}

func TrackShipment(trackingNumber string) (*models.Shipment, error) {
	return GetShipmentByTrackingNumber(trackingNumber)
}
func EditShipment(id uint, shipment *models.Shipment) error {
	return UpdateShipment(id, shipment)
}

func RemoveShipment(id uint) error {
	return DeleteShipment(id)
}
