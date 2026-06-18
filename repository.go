package supplier

import (
	"github.com/Bhandavya345/supply-chain-system/internal/database"
	"github.com/Bhandavya345/supply-chain-system/models"
)

func CreateSupplier(supplier *models.Supplier) error {
	return database.DB.Create(supplier).Error
}

func GetAllSuppliers() ([]models.Supplier, error) {

	var suppliers []models.Supplier

	err := database.DB.Find(&suppliers).Error

	return suppliers, err
}

func GetSupplierByID(id uint) (*models.Supplier, error) {

	var supplier models.Supplier

	err := database.DB.First(&supplier, id).Error

	if err != nil {
		return nil, err
	}

	return &supplier, nil
}
func UpdateSupplier(id uint, supplier *models.Supplier) error {
	return database.DB.Model(&models.Supplier{}).
		Where("id = ?", id).
		Updates(supplier).Error
}

func DeleteSupplier(id uint) error {
	return database.DB.Delete(&models.Supplier{}, id).Error
}
