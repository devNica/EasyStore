package responses

type UserAccountLoginResponseModel struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}
