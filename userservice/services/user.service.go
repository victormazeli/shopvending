package services

import (
	"github.com/sirupsen/logrus"
	"userservice/database"
	"userservice/dto"
	"userservice/models"
)

type UserService struct{}

//func (u UserService) FindUserById(id int) *models.User {
//	user := u.User
//	if err := database.Instance.Where("id = ?", id).First(&user); err != nil {
//		logrus.Errorf("Can not find user with id %d", id)
//		return nil
//
//	}
//	return user
//
//}

func (u UserService) FindUserByEmail(email string) *models.User {
	var user models.User
	if err := database.Instance.Where("email = ?", email).First(&user); err.Error != nil {
		logrus.Errorf("Can not find user with email %s", err.Error)
		return nil

	}
	return &user

}

func (u UserService) CreateUser(newUser dto.CreateNewUserDTO) *models.User {
	createNewUser := models.User{
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	createNewUser.HashPassword(createNewUser.Password)
	err := database.Instance.Create(&createNewUser)
	if err.Error != nil {
		logrus.Errorf("Error saving food to the database: %s", err.Error)
		return nil
	}

	return &createNewUser

}
