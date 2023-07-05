package repositories

import "github.com/devnica/EasyStore/entities"

type UserAccountRepository interface {
	CreateUser(newUser entities.UserAccount, rolId uint8) error
	FindUserByEmail(email string) (entities.UserAccount, error)
	FetchRolesByUserId(userId string) ([]entities.Rol, error)
}
