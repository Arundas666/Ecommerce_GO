package handlers

import (
	"firstpro/utils/response"
	"image"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

func CropImage(c *gin.Context) {

	inputImage, err := imaging.Open("./images/arundas.jpg")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open image"})
		return
	}

	
	cropRect := image.Rect(100, 100, 400, 400) // (x0, y0, x1, y1)

	
	croppedImage := imaging.Crop(inputImage, cropRect)


	err = imaging.Save(croppedImage, "./images/arundas.jpg")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save image"})
		return
		
	}

	c.JSON(200, response.ClientResponse(200, "Image cropped and saved successfully", nil, nil))

}
