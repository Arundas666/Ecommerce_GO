package usecase

import (
	"firstpro/repository"
	"firstpro/utils/models"
)

func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {

	fullOrderDetails, err := repository.GetOrderDetails(userId, page, count)
	if err != nil {
return []models.FullOrderDetails{},err
	}
	return fullOrderDetails,nil

}
