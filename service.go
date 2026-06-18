package inventory

import "github.com/Bhandavya345/supply-chain-system/models"

func CreateInventoryService(item *models.Inventory) error {
	return CreateInventory(item)
}

func GetInventoryService() ([]models.Inventory, error) {
	return GetAllInventory()
}

func EditInventory(id uint, item *models.Inventory) error {
	return UpdateInventory(id, item)
}

func RemoveInventory(id uint) error {
	return DeleteInventory(id)
}
