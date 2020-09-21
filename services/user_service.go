package services

import (
	"log"
	"go-gin-jwt-authorization-example/configs"
	"go-gin-jwt-authorization-example/dtos/dto_request"
	"go-gin-jwt-authorization-example/models"
	"go-gin-jwt-authorization-example/pkg/enums"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(userRegistrationDto dto_request.UserRegistrationDto) (err error)  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistrationDto.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}

	user := models.User{Name: userRegistrationDto.Name, Email: userRegistrationDto.Email, Password: string(hashedPassword), RoleId: enums.Student}
	if err = configs.DB.Create(&user).Error; err != nil {
		log.Println(err)
		return err;
	}

	return nil
}

func FindUserByEmail(user *models.User, email string) (err error)  {
	if err = configs.DB.Where("email = ?", email).First(user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindUserById(user *models.User, id int) (err error)  {
	if err = configs.DB.Preload("Role").Where("id = ?", id).First(user).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
