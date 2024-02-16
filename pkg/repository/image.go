package repository

import database "firstpro/pkg/db"

func GetImageUrl(productImageID int) (string, error) {
	var imageUrl string
	if err := database.DB.Raw("select product_image_url from product_images where id = ?", productImageID).Scan(&imageUrl).Error; err != nil {
		return "", err
	}
	return imageUrl, nil

}
