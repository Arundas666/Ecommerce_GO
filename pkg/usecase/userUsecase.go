package usecase

import (
	"context"
	"errors"
	"firstpro/pkg/repository"
	helper "firstpro/pkg/services"
	"firstpro/pkg/utils/models"

	"fmt"
	"net/mail"
	"strconv"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func UserSignup(user models.SignupDetail) (*models.TokenUser, error) {

	usersmatchingemail, err := repository.CheckUserExistsByEmail(user.Email)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if usersmatchingemail != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}

	usersmatchingphone, err := repository.CheckUserExistsByPhone(user.Phone)

	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if usersmatchingphone != nil {
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

	// create referral code for the user and send in details of referred id of user if it exist
	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	userReferral := str[:8]
	err = repository.CreateReferralEntry(userData, userReferral)
	if err != nil {
		return &models.TokenUser{}, err
	}

	if user.ReferralCode != "" {
		// first check whether if a user with that referralCode exist
		referredUserId, err := repository.GetUserIdFromReferrals(user.ReferralCode)
		fmt.Println("ehy reffered user id is 👺", referredUserId)
		if err != nil {
			return &models.TokenUser{}, err
		}
		if referredUserId != 0 {
			referralAmount := 100
			err := repository.UpdateReferralAmount(float64(referralAmount), referredUserId, userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
		}
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

func ApplyReferral(userID int) (string, error) {

	exist, err := repository.CartExist(userID)

	if err != nil {
		return "", err
	}

	if !exist {
		return "", errors.New("cart does not exist, can't apply offer")
	}

	referralAmount, totalCartAmount, err := repository.GetReferralAndTotalAmount(userID)
	if err != nil {
		return "", err
	}

	if totalCartAmount > referralAmount {
		totalCartAmount = totalCartAmount - referralAmount
		referralAmount = 0
	} else {
		referralAmount = referralAmount - totalCartAmount
		totalCartAmount = 0
	}

	err = repository.UpdateSomethingBasedOnUserID("referrals", "referral_amount", referralAmount, userID)
	if err != nil {
		return "", err
	}

	err = repository.UpdateSomethingBasedOnUserID("carts", "total_price", totalCartAmount, userID)
	if err != nil {
		return "", err
	}

	return "", nil
}
