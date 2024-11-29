package repository

import (
	"gin_aws/config"
	"gin_aws/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func DeleteUserByID(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
