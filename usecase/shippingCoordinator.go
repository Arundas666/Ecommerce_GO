package usecase

import (
	"firstpro/domain"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func ShippingCoordinatorLogin(shippingCoodrinatorDetails models.ShippingCoordinatorLogin) (domain.TokenShippingCoordinator, error) {

	// getting details of the admin based on the email provided
	shippingCoordinatorCompareDetails, err := repository.ShippingCoordinatorLogin(shippingCoodrinatorDetails)
	if err != nil {

		return domain.TokenShippingCoordinator{}, err
	}

	// compare password from database and that provided from admins

	err = bcrypt.CompareHashAndPassword([]byte(shippingCoordinatorCompareDetails.Password), []byte(shippingCoodrinatorDetails.Password))
	if err != nil {

		return domain.TokenShippingCoordinator{}, err
	}

	var shippingCoordinatorDetailsResponse models.ShippingCoordinatorDetailsResponse

	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&shippingCoordinatorDetailsResponse, &shippingCoordinatorCompareDetails)
	if err != nil {
		return domain.TokenShippingCoordinator{}, err
	}

	tokenString, err := helper.GenerateTokenShippingCoordinator(shippingCoordinatorDetailsResponse)

	if err != nil {
		return domain.TokenShippingCoordinator{}, err
	}

	return domain.TokenShippingCoordinator{
		ShippingCoordinator: shippingCoordinatorDetailsResponse,
		Token:               tokenString,
	}, nil

}

func UpdateShipmentStatus(shipmentStatus string,orderId string)error{
	err:=repository.UpdateShipmentStatus(orderId,shipmentStatus)
	if err!=nil{
		return  err
	}
	return nil


}