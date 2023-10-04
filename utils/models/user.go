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
	Token string
}
type LoginDetail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
