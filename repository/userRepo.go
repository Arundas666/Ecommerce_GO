package repository

import (
	"errors"
	database "firstpro/db"
	"firstpro/domain"
	"firstpro/utils/models"
	"fmt"

	"gorm.io/gorm"
)

// check whether the user is already present in the database . If there recommend to login
func CheckUserAvailability(email string) bool {

	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := database.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func CheckUserAvailabilityWithUserID(userId int) bool {

	var count int

	if err := database.DB.Raw("select count(*) from users where id=?", userId).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0

}
func CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := database.DB.Where(&domain.User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	result := database.DB.Where(&domain.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func UserSignup(user models.SignupDetail) (models.SignupDetailResponse, error) {
	var signupDetail models.SignupDetailResponse
	err := database.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,email,phone", user.FirstName, user.LastName, user.Email, user.Password, user.Phone).Scan(&signupDetail).Error
	if err != nil {
		fmt.Println("Repository error:", err)
		return models.SignupDetailResponse{}, err
	}
	return signupDetail, nil

}

func FindUserDetailsByEmail(user models.LoginDetail) (models.UserLoginResponse, error) {
	var userdetails models.UserLoginResponse

	err := database.DB.Raw(
		`SELECT * FROM users where email = ? and blocked = false`, user.Email).Scan(&userdetails).Error

	if err != nil {
		return models.UserLoginResponse{}, errors.New("error checking user details")
	}
	return userdetails, nil

}

func GetAllAddress(userId int) (models.AddressInfoResponse, error) {
	var addressInfoResponse models.AddressInfoResponse
	if err := database.DB.Raw("select * from addresses where user_id = ?", userId).Scan(&addressInfoResponse).Error; err != nil {
		return models.AddressInfoResponse{}, err
	}
	fmt.Println(addressInfoResponse, "HEyy")
	return addressInfoResponse, nil
}
func GetAddressFromOrderId(orderId string) (models.AddressInfoResponse, error) {
	var addressInfoResponse models.AddressInfoResponse
	var addressId int
	if err := database.DB.Raw("select address_id from orders where order_id =?", orderId).Scan(&addressId).Error; err != nil {
		return models.AddressInfoResponse{}, errors.New("first in orders")
	}
	if err := database.DB.Raw("select * from addresses where id=?", addressId).Scan(&addressInfoResponse).Error; err != nil {
		return models.AddressInfoResponse{}, errors.New("second  in address")
	}
	return addressInfoResponse, nil
}

func AddAddress(userId int, address models.AddressInfo) error {

	if err := database.DB.Raw("insert into addresses(user_id,name,house_name,street,city,state,pin)  values(?,?,?,?,?,?,?)", userId, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Scan(&address).Error; err != nil {
		return err
	}
	return nil
}
func UserDetails(userID int) (models.UsersProfileDetails, error) {

	var userDetails models.UsersProfileDetails
	err := database.DB.Raw("select users.firstname,users.lastname,users.email,users.phone from users  where users.id = ?", userID).Row().Scan(&userDetails.Firstname, &userDetails.Lastname, &userDetails.Email, &userDetails.Phone)
	if err != nil {
		return models.UsersProfileDetails{}, err
	}
	return userDetails, nil
}

func UpdateUserEmail(email string, userID int) error {

	err := database.DB.Exec("update users set email = ? where id = ?", email, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateUserPhone(phone string, userID int) error {

	err := database.DB.Exec("update users set phone = ? where id = ?", phone, userID).Error
	if err != nil {
		return err
	}
	return nil

}

func UpdateFirstName(name string, userID int) error {

	err := database.DB.Exec("update users set firstname = ? where id = ?", name, userID).Error
	if err != nil {
		return err
	}
	return nil

}
func UpdateLastName(name string, userID int) error {

	err := database.DB.Exec("update users set lastname = ? where id = ?", name, userID).Error
	if err != nil {
		return err
	}
	return nil

}
func UserPassword(userID int) (string, error) {

	var userPassword string
	err := database.DB.Raw("select password from users where id = ?", userID).Scan(&userPassword).Error
	if err != nil {
		return "", err
	}
	return userPassword, nil

}
func UpdateUserPassword(password string, userID int) error {
	err := database.DB.Exec("update users set password = ? where id = ?", password, userID).Error
	if err != nil {
		return err
	}
	fmt.Println("password Updated succesfully")
	return nil
}

func GetAllAddresses(userID int) ([]models.AddressInfoResponse, error) {

	var addressResponse []models.AddressInfoResponse
	err := database.DB.Raw(`select * from addresses where user_id = $1`, userID).Scan(&addressResponse).Error
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}

	return addressResponse, nil

}

func GetAllPaymentOption() ([]models.PaymentDetails, error) {

	var paymentMethods []models.PaymentDetails
	err := database.DB.Raw("select * from payment_methods").Scan(&paymentMethods).Error
	if err != nil {
		return []models.PaymentDetails{}, err
	}

	return paymentMethods, nil

}
