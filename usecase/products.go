package usecase

import (
	"firstpro/repository"
	"firstpro/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBrief, error) {

	productsBrief, err := repository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBrief{}, err
	}
	for i := range productsBrief {
		p := &productsBrief[i]
		if p.Quantity == 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	return productsBrief, nil

}

func FilterCategory(data map[string]int) ([]models.ProductBrief, error) {

	err := repository.CheckValidityOfCategory(data)
	if err != nil {
		return []models.ProductBrief{}, err
	}

	var productFromCategory []models.ProductBrief
	for _, id := range data {

		product, err := repository.GetProductFromCategory(id)
		if err != nil {
			return []models.ProductBrief{}, err
		}
		for _, product := range product {

			quantity, err := repository.GetQuantityFromProductID(product.ID)
			if err != nil {
				return []models.ProductBrief{}, err
			}
			if quantity == 0 {
				product.ProductStatus = "out of stock"
			} else {
				product.ProductStatus = "in stock"
			}
			if product.ID != 0 {
				productFromCategory = append(productFromCategory, product)
			}
		}

		// if a product exist for that genre. Then only append it

	}
	return productFromCategory, nil

}
