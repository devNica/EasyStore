package impl

import (
	"context"
	"time"

	"github.com/devnica/EasyStore/commons"
	security "github.com/devnica/EasyStore/commons/security"
	utils "github.com/devnica/EasyStore/commons/utils"
	"github.com/devnica/EasyStore/configurations"
	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/devnica/EasyStore/models/request"
	"github.com/devnica/EasyStore/models/response"
	"github.com/devnica/EasyStore/repositories"
	"github.com/devnica/EasyStore/services"
	"github.com/google/uuid"
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
	newUser request.UserAccountRegisterRequestModel) {

	accountStatus := commons.GetAccountStatusFromDictionary()
	statusId := commons.GetKeyId("unverifiableIdentity", accountStatus)

	hash := security.GeneratePasswordHash(newUser.Password, &srv.Argon2Config)

	user := dto.UserRegisterDTOModel{
		Id:            uuid.New(),
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

func (srv *userAccountServiceImpl) GetUserByEmail(ctx context.Context, data request.UserAccountLoginRequestModel) response.UserAccountLoginResponseModel {

	user, err := srv.UserAccountRepository.FindUserByEmail(data.Email)
	exceptions.PanicLogging(err)

	match, matchErr := security.ComparePasswordAndHash(data.Password, user.Password, &srv.Argon2Config)

	if match != true {
		panic(exceptions.BadRequestError{
			Message: matchErr.Error(),
		})
	}

	roles, err := srv.UserAccountRepository.FetchRolesByUserId(user.Id.String())

	login := response.UserAccountLoginResponseModel{
		UserId: user.Id.String(),
		Email:  user.Email,
	}

	rolesMap := utils.ConvertRolesToMaps(roles)

	login.Token = security.GenerateToken(login.UserId, rolesMap)

	return login
}
