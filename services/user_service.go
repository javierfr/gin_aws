package services

import (
	"gin_aws/models"
	"gin_aws/repository"
)

func GetUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func CreateUser(user *models.User) error {
	return repository.CreateUser(user)
}

func GetUser(id uint) (*models.User, error) {
	return repository.GetUserByID(id)
}

func DeleteUser(id uint) error {
	return repository.DeleteUserByID(id)
}
