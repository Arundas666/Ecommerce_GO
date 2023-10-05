package usecase

import (
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

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminCompareDetails.Password), bcrypt.DefaultCost)

	// compare password from database and that provided from admins

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(adminDetails.Password))
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

	productDetails, err := repository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	return models.CompleteAdminDashboard{

		DashboardUser:    userDetails,
		DashBoardProduct: productDetails,
	}, nil

}
