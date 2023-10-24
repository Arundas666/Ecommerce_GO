package usecase

import (
	"context"
	"errors"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"
	"net/mail"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(user models.SignupDetail) (*models.TokenUser, error) {
	fmt.Println(user, "👏")
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return &models.TokenUser{}, errors.New("invalid email format")
	}

	// Phone number validation
	if len(user.Phone) != 10 {
		return &models.TokenUser{}, errors.New("phone number should have 10 digits")
	}
	//check whether the user already exsist by looking the email and the phone number provided
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email, "🙌")
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}

	phone, err := repository.CheckUserExistsByPhone(user.Phone)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}

	//if the signing up is a new user then hashing the password
	hashedPassword, err := helper.PasswordHashing(user.Password)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in hashing password")
	}

	user.Password = hashedPassword
	//after hashing adding the user detail into the database and taking the added user detail to the userdata
	userData, err := repository.UserSignup(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add the user ")
	}

	//creating a jwt token for the new user with the detail that has been stored in the database
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("counldnt create access token due to error")
	}

	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("counldnt create refresh token due to error")
	}

	return &models.TokenUser{
		Users:        userData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func UserLoginWithPassword(user models.LoginDetail) (*models.TokenUser, error) {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return &models.TokenUser{}, errors.New("invalid email format")
	}
	email, err := repository.CheckUserExistsByEmail(user.Email)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return &models.TokenUser{}, errors.New("email  does not exsist")
	}

	userDetails, err := repository.FindUserDetailsByEmail(user)
	if err != nil {
		return &models.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(user.Password))

	if err != nil {
		return &models.TokenUser{}, errors.New("password not matching")
	}
	var user_details models.SignupDetailResponse
	err = copier.Copy(&user_details, &userDetails)
	if err != nil {
		return &models.TokenUser{}, err
	}
	accessToken, err := helper.GenerateAccessToken(user_details)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(user_details)
	if err != nil {
		return &models.TokenUser{}, errors.New("counldnt create refresh token due to error")
	}

	return &models.TokenUser{
		Users:        user_details,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GetAllAddress(userId int) (models.AddressInfoResponse, error) {
	addressInfo, err := repository.GetAllAddress(userId)
	if err != nil {
		return models.AddressInfoResponse{}, err
	}
	return addressInfo, nil

}
func AddAddress(userId int, address models.AddressInfo) error {
	if err := repository.AddAddress(userId, address); err != nil {
		return err
	}

	return nil

}

func UserDetails(userID int) (models.UsersProfileDetails, error) {
	return repository.UserDetails(userID)
}

func UpdateUserDetails(userDetails models.UsersProfileDetails, userID int) (models.UsersProfileDetails, error) {
	userExist := repository.CheckUserAvailabilityWithUserID(userID)

	if !userExist {
		return models.UsersProfileDetails{}, errors.New("user doesnt exist")
	}
	// which all field are not empty (which are provided from the front end should be updated)
	if userDetails.Email != "" {
		repository.UpdateUserEmail(userDetails.Email, userID)
	}
	if userDetails.Firstname != "" {
		repository.UpdateFirstName(userDetails.Firstname, userID)

	}
	if userDetails.Firstname != "" {
		repository.UpdateLastName(userDetails.Lastname, userID)
	}

	if userDetails.Phone != "" {
		repository.UpdateUserPhone(userDetails.Phone, userID)
	}
	return repository.UserDetails(userID)
}
func UpdatePassword(ctx context.Context, body models.UpdatePassword) error {
	var userID int
	var ok bool
	if userID, ok = ctx.Value("userID").(int); !ok {
		return errors.New("error retrieving user details")
	}
	fmt.Println("user id is", userID)
	userPassword, err := repository.UserPassword(userID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(body.OldPassword))
	if err != nil {
		return errors.New("password incorrect")
	}
	if body.NewPassword != body.ConfirmNewPassword {
		return errors.New("password not matching")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)
	if err != nil {
		return err
	}
	if err := repository.UpdateUserPassword(string(hashedPassword), userID); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Checkout(userID int) (models.CheckoutDetails, error) {

	// list all address added by the user
	allUserAddress, err := repository.GetAllAddresses(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get available payment options
	paymentDetails, err := repository.GetAllPaymentOption()
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get all items from users cart
	cartItems, err := repository.DisplayCart(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get grand total of all the product
	grandTotal, err := repository.GetTotalPrice(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		Payment_Method:      paymentDetails,
		Cart:                cartItems,

		Grand_Total: grandTotal.TotalPrice,
		Total_Price: grandTotal.FinalPrice,
	}, nil
}
