package repository

import (
	"errors"
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

func CheckValidityOfCategory(data map[string]int) error {

	for _, id := range data {
		var count int
		err := database.DB.Raw("select count(*) from categories where id = ?", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count < 1 {
			return errors.New("genre does not exist")
		}
	}
	return nil
}
func GetProductFromCategory(id int) (models.ProductBrief, error) {

	var product models.ProductBrief
	err := database.DB.Raw(`
		SELECT *
		FROM products
		JOIN categories ON products.category_id = categories.id
		 where categories.id = ?
	`, id).Scan(&product).Error

	if err != nil {
		return models.ProductBrief{}, err
	}

	return product, nil

}

func GetQuantityFromProductID(id int) (int, error) {

	var quantity int
	err := database.DB.Raw("select quantity from products where id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	return quantity, nil

}
