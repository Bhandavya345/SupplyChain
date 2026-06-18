package supplier

import "github.com/Bhandavya345/supply-chain-system/models"

func AddSupplier(supplier *models.Supplier) error {
	return CreateSupplier(supplier)
}

func FetchAllSuppliers() ([]models.Supplier, error) {
	return GetAllSuppliers()
}

func FetchSupplierByID(id uint) (*models.Supplier, error) {
	return GetSupplierByID(id)
}
func EditSupplier(id uint, supplier *models.Supplier) error {
	return UpdateSupplier(id, supplier)
}

func RemoveSupplier(id uint) error {
	return DeleteSupplier(id)
}
