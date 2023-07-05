package request

type UserAccountRegisterRequestModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserAccountLoginRequestModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatePersonalInfoRequestModel struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	DNI         string `json:"dni" valiate:"required"`
	FirstName   string `json:"firstname" validate:"required"`
	LastName    string `json:"lastname" validate:"required"`
	Address     string `json:"address" validate:"required"`
	BirthDate   string `json:"birthdate" validate:"required"`
}
