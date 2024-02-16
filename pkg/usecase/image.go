package usecase

import (
	"firstpro/pkg/repository"
	"fmt"
)

func CropImage(productImageId int) (string, error) {

	imageUrl, err := repository.GetImageUrl(productImageId)
	fmt.Println("image url is ", imageUrl)
	if err != nil {
		return "", err
	}
	return imageUrl, nil

}
