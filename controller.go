package auth

import (
	"net/http"

	"github.com/Bhandavya345/supply-chain-system/internal/utils"
	"github.com/Bhandavya345/supply-chain-system/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register User
// @Description Register a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} map[string]interface{}
// @Router /api/auth/register [post]
func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := RegisterUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

// Login godoc
// @Summary Login User
// @Description Login and generate JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/login [post]
func Login(c *gin.Context) {

	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := GetUserByEmail(loginRequest.Email)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginRequest.Password),
	)

	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})

		return
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
