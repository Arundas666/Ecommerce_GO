package usecase

import (
	"errors"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"
	"fmt"
)

func UserSignup(user models.SignupDetail) (*models.TokenUser, error) {
	fmt.Println(user, "👏")
	//check whether the user already exsist by looking the email and the phone number provided
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email, "🙌")
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}
	fmt.Println("fghfvwhgbfyh")
	phone, err := repository.CheckUserExistsByPhone(user.Phone)
	fmt.Println(phone, "ewrtty")
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
	tokenString, err := helper.GenerateTokenClients(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not create token due to some internal error")
	}
	return &models.TokenUser{
		Users: userData,
		Token: tokenString,
	}, nil

}


// func UserLogin(user models.LoginDetail)(*models.TokenUser,error) {
// 	email,err:=repository.CheckUserExistsByEmail(user.Email)
// 	if err!=nil{
// 		return &models.TokenUser{}, errors.New("error with server")
// 	}
// 	if email==nil{
// 		return &models.TokenUser{}, errors.New("email with this user does not exsist")
// 	}
// 	phone,err:=repository.CheckUserExistsByPhone(user.Password)
// 	if err!=nil{
// 		return &models.TokenUser{}, errors.New("error with server")
// 	}
// 	if phone==nil{
// 		return &models.TokenUser{}, errors.New("password is incorrect")
// 	}

// 	userlog,err:=repository.UserLogin(user)

// }



