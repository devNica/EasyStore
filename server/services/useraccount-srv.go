package services

import (
	"context"

	"github.com/devnica/EasyStore/models/request"
	"github.com/devnica/EasyStore/models/response"
)

type UserAccountService interface {
	UserRegister(ctx context.Context, newUser request.UserAccountRegisterRequestModel)
	GetUserByEmail(ctx context.Context, user request.UserAccountLoginRequestModel) response.UserAccountLoginResponseModel
	UpdatePersonalInfo(ctx context.Context, personalInfo request.UpdatePersonalInfoRequestModel, userId string)
}
