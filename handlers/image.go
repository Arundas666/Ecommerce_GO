package handlers

import (
	"firstpro/usecase"
	"firstpro/utils/response"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func CropImage(c *gin.Context) {
	imageId := c.Query("product_image_id")
	imageID, err := strconv.Atoi(imageId)
	if err != nil {
		errRes := response.ClientResponse(500, "error in string conversion", nil, err)
		c.JSON(500, errRes)
		return
	}
	imageUrl, err := usecase.CropImage(imageID)
	if err != nil {
		errRes := response.ClientResponse(500, "error in cropping", nil, err)
		c.JSON(500, errRes)
		return
	}

	inputImage, err := imaging.Open(imageUrl)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open image"})
		return
	}

	cropRect := image.Rect(100, 100, 400, 400) // (x0, y0, x1, y1)

	croppedImage := imaging.Crop(inputImage, cropRect)

	err = imaging.Save(croppedImage, imageUrl)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save image"})
		return

	}
	
	
	//VESRILE NEXTJS FIREBASE GO
	
	
	c.JSON(200, response.ClientResponse(200, "Image cropped and saved successfully", nil, nil))

}
