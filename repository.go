package auth

import (
	"github.com/Bhandavya345/supply-chain-system/internal/database"
	"github.com/Bhandavya345/supply-chain-system/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func FindUserByEmail(email string) (*models.User, error) {

	var user models.User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
