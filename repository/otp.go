package repository

import (
	database "firstpro/db"
	"firstpro/utils/models"
)


func FindUserByMobileNumber(phone string)bool{
	var count int
	if err := database.DB.Raw("select count(*) from users where phone = ?", phone).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}
func UserDetailsUsingPhone(phone string) (models.SignupDetailResponse, error) {

	var usersDetails models.SignupDetailResponse
	if err := database.DB.Raw("select * from users where phone = ?", phone).Scan(&usersDetails).Error; err != nil {
		return models.SignupDetailResponse{}, err
	}

	return usersDetails, nil

}