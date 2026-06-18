package tests

import (
	"testing"

	"github.com/Bhandavya345/supply-chain-system/models"

	"github.com/stretchr/testify/assert"
)

func TestInventoryCreation(t *testing.T) {

	item := models.Inventory{
		ProductName: "Laptop",
		Quantity:    10,
		//Price:       50000,
	}

	assert.Equal(t, "Laptop", item.ProductName)
	assert.Equal(t, 10, item.Quantity)
}
