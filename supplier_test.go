package tests

import (
	"testing"

	"github.com/Bhandavya345/supply-chain-system/models"

	"github.com/stretchr/testify/assert"
)

func TestSupplierCreation(t *testing.T) {

	supplier := models.Supplier{
		SupplierName: "ABC Traders",
		Email:        "abc@gmail.com",
	}

	assert.Equal(t, "ABC Traders", supplier.SupplierName)
	assert.Equal(t, "abc@gmail.com", supplier.Email)
}
