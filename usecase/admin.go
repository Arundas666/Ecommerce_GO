package usecase

import (
	"errors"
	"firstpro/domain"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := repository.LoginHandler(adminDetails)
	if err != nil {

		return domain.TokenAdmin{}, err
	}

	// compare password from database and that provided from admins

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		fmt.Println("👺", err)
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	//  copy all details except password and sent it back to the front end
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil

}

func DashBoard() (models.CompleteAdminDashboard, error) {

	userDetails, err := repository.DashboardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := repository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	productDetails, err := repository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	return models.CompleteAdminDashboard{

		DashboardUser:    userDetails,
		DashBoardProduct: productDetails,
		DashboardOrder:   orderDetails,
	}, nil

}
func ApproveOrder(orderID string) error {
	ok, err := repository.CheckOrderID(orderID)
	if !ok {
		return err
	}

	shipmentStatus, err := repository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}

	if shipmentStatus == "cancelled" {

		return errors.New("the order is cancelled, cannot approve it")
	}

	if shipmentStatus == "pending" {

		return errors.New("the order is pending, cannot approve it")
	}
	if shipmentStatus == "processing" {
		fmt.Println("reached here")
		err := repository.ApproveOrder(orderID)

		if err != nil {
			return err
		}

		return nil
	}

	// if the shipment status is not processing or cancelled. Then it is defenetely cancelled
	return nil

}
