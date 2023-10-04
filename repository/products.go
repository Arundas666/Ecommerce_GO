package repository

import (
	database "firstpro/db"
	"firstpro/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productsBrief []models.ProductBrief
	err := database.DB.Raw(`
		SELECT * FROM products limit ? offset ?
	`, count, offset).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	return productsBrief, nil

}
