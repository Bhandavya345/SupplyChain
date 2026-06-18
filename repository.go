package inventory

import (
	"github.com/Bhandavya345/supply-chain-system/internal/database"
	"github.com/Bhandavya345/supply-chain-system/models"
)

func CreateInventory(item *models.Inventory) error {
	return database.DB.Create(item).Error
}

func GetAllInventory() ([]models.Inventory, error) {

	var items []models.Inventory

	err := database.DB.Find(&items).Error

	return items, err
}

func GetInventoryByID(id uint) (*models.Inventory, error) {

	var item models.Inventory

	err := database.DB.First(&item, id).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func UpdateInventory(id uint, item *models.Inventory) error {
	return database.DB.Model(&models.Inventory{}).
		Where("id = ?", id).
		Updates(item).Error
}

func DeleteInventory(id uint) error {
	return database.DB.Delete(&models.Inventory{}, id).Error
}
