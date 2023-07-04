package repositories

import "github.com/devnica/EasyStore/entities"

type UserAccountRepository interface {
	CreateUser(newUser entities.UserAccount, rolId uint8) error
}
