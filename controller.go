package supplier

import (
	"net/http"
	"strconv"

	"github.com/Bhandavya345/supply-chain-system/models"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create Supplier
// @Tags Suppliers
// @Accept json
// @Param supplier body models.Supplier true "Supplier"
// @Router /api/suppliers [post]
func Create(c *gin.Context) {

	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := AddSupplier(&supplier)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Supplier created successfully",
		"data":    supplier,
	})
}

// GetAll godoc
// @Summary Get All Suppliers
// @Tags Suppliers
// @Produce json
// @Success 200 {array} models.Supplier
// @Router /api/suppliers [get]
func GetAll(c *gin.Context) {

	suppliers, err := FetchAllSuppliers()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

// GetSupplierByID godoc
// @Summary Get Supplier by ID
// @Description Get supplier details by ID
// @Tags Suppliers
// @Produce json
// @Param id path int true "Supplier ID"
// @Success 200 {object} models.Supplier
// @Failure 404 {object} map[string]interface{}
// @Router /api/suppliers/{id} [get]
func GetByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid supplier id",
		})
		return
	}

	supplier, err := FetchSupplierByID(uint(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Supplier not found",
		})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// Update godoc
// @Summary Update Supplier
// @Tags Suppliers
// @Param id path int true "Supplier ID"
// @Param supplier body models.Supplier true "Supplier"
// @Router /api/suppliers/{id} [put]
func Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := EditSupplier(uint(id), &supplier)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Supplier updated"})
}

// Delete godoc
// @Summary Delete Supplier
// @Tags Suppliers
// @Param id path int true "Supplier ID"
// @Router /api/suppliers/{id} [delete]
func Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := RemoveSupplier(uint(id))

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Supplier deleted"})
}
