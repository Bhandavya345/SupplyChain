package inventory

import (
	"net/http"
	"strconv"

	"github.com/Bhandavya345/supply-chain-system/models"
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create Inventory
// @Tags Inventory
// @Accept json
// @Produce json
// @Param inventory body models.Inventory true "Inventory"
// @Success 201 {object} models.Inventory
// @Router /api/inventory [post]
func Create(c *gin.Context) {

	var item models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := CreateInventoryService(&item)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Inventory created successfully",
	})
}

// GetAll godoc
// @Summary Get All Inventory
// @Tags Inventory
// @Produce json
// @Success 200 {array} models.Inventory
// @Router /api/inventory [get]
func GetAll(c *gin.Context) {

	items, err := GetInventoryService()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, items)
}

// GetInventoryByID godoc
// @Summary Get Inventory by ID
// @Description Get inventory details by ID
// @Tags Inventory
// @Produce json
// @Param id path int true "Inventory ID"
// @Success 200 {object} models.Inventory
// @Failure 404 {object} map[string]interface{}
// @Router /api/inventory/{id} [get]
func GetByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid inventory id",
		})
		return
	}

	inventory, err := GetInventoryByID(uint(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Inventory not found",
		})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// Update godoc
// @Summary Update Inventory
// @Tags Inventory
// @Accept json
// @Param id path int true "Inventory ID"
// @Param inventory body models.Inventory true "Inventory"
// @Router /api/inventory/{id} [put]
func Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var item models.Inventory

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := EditInventory(uint(id), &item)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Inventory updated"})
}

// Delete godoc
// @Summary Delete Inventory
// @Tags Inventory
// @Param id path int true "Inventory ID"
// @Router /api/inventory/{id} [delete]
func Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := RemoveInventory(uint(id))

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Inventory deleted"})
}
