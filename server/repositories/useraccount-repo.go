package repositories

import (
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/google/uuid"
)

type UserAccountRepository interface {
	CreateUser(newUser dto.UserRegisterDTOModel, roleId uint8) error
	FindUserByEmail(email string) (dao.FindUserDAOModel, error)
	FetchRolesByUserId(userId string) ([]dao.RolDAOModel, error)
	InsertPersonalInfo(personalInfo dto.PersonalInfoDTOModel, userId string) error
	InsertRoleToUserAccount(userId uuid.UUID, roleId uint8) error
	FetchStatusAccountByUserId(userId uuid.UUID) (dao.AccounStatusDAOModel, error)
}
