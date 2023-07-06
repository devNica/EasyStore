package impl

import (
	"context"
	"fmt"
	"log"
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
	repositories.AdminCommitRepository
	configurations.Argon2Config
}

func NewUserAccountServiceImpl(
	userRepo *repositories.UserAccountRepository,
	adminRepo *repositories.AdminCommitRepository,
	argon *configurations.Argon2Config) services.UserAccountService {
	return &userAccountServiceImpl{
		UserAccountRepository: *userRepo,
		AdminCommitRepository: *adminRepo,
		Argon2Config:          *argon}
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
	roleId := commons.GetKeyId("owners", roles)

	err := srv.UserAccountRepository.CreateUser(user, roleId)
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

	fmt.Println(login.Token)

	return login
}

func (srv *userAccountServiceImpl) UpdatePersonalInfo(ctx context.Context, data request.UpdatePersonalInfoRequestModel, userId string) {

	personalInfo := dto.PersonalInfoDTOModel{
		PhoneNumber:   data.PhoneNumber,
		DNI:           data.DNI,
		FirstName:     data.FirstName,
		LastName:      data.LastName,
		Address:       data.Address,
		BirthDate:     data.BirthDate,
		TwoFactorAuth: true,
		UpdatedAt:     time.Now(),
		StatusId:      2,
	}

	err := srv.UserAccountRepository.InsertPersonalInfo(personalInfo, userId)

	if err != nil {
		log.Fatalf(err.Error())
	}

	kycReview := dto.KYCReviewDTOModel{
		Id:             uuid.New(),
		UserRef:        userId,
		PreRevStatus:   "awaitingReview",
		CreatedAt:      time.Now(),
		ReviewStatusId: 1,
	}

	err = srv.AdminCommitRepository.InsertKYCReviewHistory(kycReview)

	if err != nil {
		log.Fatalf(err.Error())
	}

}
