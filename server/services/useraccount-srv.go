package services

import (
	"context"

	"github.com/devnica/EasyStore/dto/requests"
)

type UserAccountService interface {
	UserRegister(ctx context.Context, newUser requests.UserAccountRegisterRequestModel)
}
