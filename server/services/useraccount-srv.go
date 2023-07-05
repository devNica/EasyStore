package services

import (
	"context"

	"github.com/devnica/EasyStore/dto/requests"
	"github.com/devnica/EasyStore/dto/responses"
)

type UserAccountService interface {
	UserRegister(ctx context.Context, newUser requests.UserAccountRegisterRequestModel)
	GetUserByEmail(ctx context.Context, user requests.UserAccountLoginRequestModel) responses.UserAccountLoginResponseModel
}
