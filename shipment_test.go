package tests

import (
	"testing"

	"github.com/Bhandavya345/supply-chain-system/models"

	"github.com/stretchr/testify/assert"
)

func TestShipmentCreation(t *testing.T) {

	shipment := models.Shipment{
		TrackingNumber: "TRK1001",
		Status:         "In Transit",
	}

	assert.Equal(t, "TRK1001", shipment.TrackingNumber)
	assert.Equal(t, "In Transit", shipment.Status)
}
