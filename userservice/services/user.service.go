package services

import (
	"github.com/sirupsen/logrus"
	"userservice/database"
	"userservice/dto"
	"userservice/models"
)

type UserService struct{}

func (u UserService) FindUserById(id int) *models.User {
	var user models.User
	result := database.Instance.Where("id = ?", id).First(&user)

	if result.Error != nil {
		logrus.Errorf("Can not find user with id %d", id)
		return nil
	}
	return &user

}

func (u UserService) FindUserByEmail(email string) *models.User {
	var user models.User
	result := database.Instance.Where("email = ?", email).First(&user)

	if result.Error != nil {
		logrus.Errorf("Can not find user with email %s", result.Error)
		return nil
	}
	return &user

}

func (u UserService) FindUserOTPCode(code string) *models.User {
	var user models.User
	result := database.Instance.Where("otp_code = ?", code).First(&user)

	if result.Error != nil {
		logrus.Errorf("Can not find code %s", result.Error)
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
	result := database.Instance.Create(&createNewUser)
	if result.Error != nil {
		logrus.Errorf("Error saving user to the database: %s", result.Error)
		return nil
	}

	return &createNewUser

}

func (u UserService) UpdateBasicDetails(userId int, data map[string]interface{}) *models.User {
	var user models.User
	result := database.Instance.Model(&user).Where("id = ?", userId).Updates(data)

	if result != nil {
		logrus.Errorf("Error saving food to the database: %s", result.Error)
		return nil
	}
	return &user

}
