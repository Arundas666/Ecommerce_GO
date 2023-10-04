package models

type SignupDetail struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type SignupDetailResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
type TokenUser struct {
	Users SignupDetailResponse
	AccessToken string
	RefreshToken string
}
type LoginDetail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct{
	Id uint `json:"id"`
	UserId uint `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`

}