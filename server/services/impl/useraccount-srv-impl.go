package impl

import (
	"context"
	"time"

	"github.com/devnica/EasyStore/commons"
	argon2 "github.com/devnica/EasyStore/commons/security"
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/dto/requests"
	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/repositories"
	"github.com/devnica/EasyStore/services"
)

type userAccountServiceImpl struct {
	repositories.UserAccountRepository
	configurations.Argon2Config
}

func NewUserAccountServiceImpl(
	repo *repositories.UserAccountRepository,
	argon *configurations.Argon2Config) services.UserAccountService {
	return &userAccountServiceImpl{UserAccountRepository: *repo, Argon2Config: *argon}
}

func (srv *userAccountServiceImpl) UserRegister(
	ctx context.Context,
	newUser requests.UserAccountRegisterRequestModel) {

	accountStatus := commons.GetAccountStatusFromDictionary()
	statusId := commons.GetKeyId("unverifiableIdentity", accountStatus)

	hash := argon2.GeneratePasswordHash(newUser.Password, &srv.Argon2Config)

	user := entities.UserAccount{
		Email:         newUser.Email,
		Password:      hash,
		TwoFactorAuth: false,
		CreatedAt:     time.Now(),
		StatusId:      statusId,
	}

	roles := commons.GetRolesFromDictionary()
	rolId := commons.GetKeyId("owners", roles)

	err := srv.UserAccountRepository.CreateUser(user, rolId)
	exceptions.PanicLogging(err)
}
