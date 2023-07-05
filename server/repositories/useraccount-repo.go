package repositories

import (
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
)

type UserAccountRepository interface {
	CreateUser(newUser dto.UserRegisterDTOModel, rolId uint8) error
	FindUserByEmail(email string) (dao.FindUserDAOModel, error)
	FetchRolesByUserId(userId string) ([]dao.RolDAOModel, error)
}
