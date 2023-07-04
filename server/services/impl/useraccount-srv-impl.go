package impl

import (
	"context"
	"time"

	"github.com/devnica/EasyStore/commons"
	"github.com/devnica/EasyStore/dto/requests"
	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/repositories"
	"github.com/devnica/EasyStore/services"
)

type userAccountServiceImpl struct {
	repositories.UserAccountRepository
}

func NewUserAccountServiceImpl(repo *repositories.UserAccountRepository) services.UserAccountService {
	return &userAccountServiceImpl{UserAccountRepository: *repo}
}

func (srv *userAccountServiceImpl) UserRegister(
	ctx context.Context,
	newUser requests.UserAccountRegisterRequestModel) {

	accountStatus := commons.GetAccountStatusFromDictionary()
	statusId := commons.GetKeyId("unverifiableIdentity", accountStatus)

	user := entities.UserAccount{
		Email:         newUser.Email,
		Password:      newUser.Password,
		TwoFactorAuth: false,
		CreatedAt:     time.Now(),
		StatusId:      statusId,
	}

	roles := commons.GetRolesFromDictionary()
	rolId := commons.GetKeyId("owners", roles)

	err := srv.UserAccountRepository.CreateUser(user, rolId)
	exceptions.PanicLogging(err)
}
