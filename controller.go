package shipment

import (
	"net/http"
	"strconv"

	"github.com/Bhandavya345/supply-chain-system/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create Shipment
// @Tags Shipments
// @Accept json
// @Param shipment body models.Shipment true "Shipment"
// @Router /api/shipments [post]
func Create(c *gin.Context) {

	var shipment models.Shipment

	if err := c.ShouldBindJSON(&shipment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := AddShipment(&shipment)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Shipment created successfully",
		"data":    shipment,
	})
}

// GetAll godoc
// @Summary Get All Shipments
// @Tags Shipments
// @Produce json
// @Success 200 {array} models.Shipment
// @Router /api/shipments [get]
func GetAll(c *gin.Context) {

	shipments, err := FetchAllShipments()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shipments)
}

// GetShipmentByID godoc
// @Summary Get Shipment by ID
// @Description Get shipment details by ID
// @Tags Shipments
// @Produce json
// @Param id path int true "Shipment ID"
// @Success 200 {object} models.Shipment
// @Failure 404 {object} map[string]interface{}
// @Router /api/shipments/{id} [get]
func GetByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid shipment id",
		})
		return
	}

	shipment, err := FetchShipmentByID(uint(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Shipment not found",
		})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

// Track godoc
// @Summary Track Shipment
// @Tags Shipments
// @Param trackingNumber path string true "Tracking Number"
// @Router /api/shipments/track/{trackingNumber} [get]
func Track(c *gin.Context) {

	trackingNumber := c.Param("trackingNumber")

	shipment, err := TrackShipment(trackingNumber)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Shipment not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tracking_number": shipment.TrackingNumber,
		"status":          shipment.Status,
		"source":          shipment.Source,
		"destination":     shipment.Destination,
	})
}

// Update godoc
// @Summary Update Shipment
// @Tags Shipments
// @Accept json
// @Param id path int true "Shipment ID"
// @Param shipment body models.Shipment true "Shipment"
// @Router /api/shipments/{id} [put]
func Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var shipment models.Shipment

	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := EditShipment(uint(id), &shipment)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Shipment updated"})
}

// Delete godoc
// @Summary Delete Shipment
// @Tags Shipments
// @Param id path int true "Shipment ID"
// @Router /api/shipments/{id} [delete]
func Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := RemoveShipment(uint(id))

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Shipment deleted"})
}
