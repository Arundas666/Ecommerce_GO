package repository

import (
	"errors"
	"firstpro/pkg/domain"
	database "firstpro/pkg/db"
	"firstpro/pkg/utils/models"
	"fmt"

	"gorm.io/gorm"
)

func FindUserByMobileNumber(phone string) (*domain.User, error) {

	// var count int
	// if err := database.DB.Raw("select count(*) from users where phone = ?", phone).Scan(&count).Error; err != nil {
	// 	return false
	// }

	// return count > 0
	var user domain.User
	result := database.DB.Where(&domain.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil

}
func UserDetailsUsingPhone(phone string) (models.SignupDetailResponse, error) {

	var usersDetails models.SignupDetailResponse
	if err := database.DB.Raw("select * from users where phone = ?", phone).Scan(&usersDetails).Error; err != nil {
		return models.SignupDetailResponse{}, err
	}

	return usersDetails, nil

}
func FindUserByEmail(email string) (bool, error) {

	var count int
	if err := database.DB.Raw("select count(*) from users where email = ?", email).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
func GetUserPhoneByEmail(email string) (string, error) {
	fmt.Println(email)
	var phone string
	if err := database.DB.Raw("select phone from users where email = ?", email).Scan(&phone).Error; err != nil {
		return "", err
	}

	return phone, nil

}
