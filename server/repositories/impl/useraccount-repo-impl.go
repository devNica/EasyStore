package impl

import (
	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userAccountRepositoryImpl struct {
	*gorm.DB
}

func NewUserAccountRepositoryImpl(DB *gorm.DB) repositories.UserAccountRepository {
	return &userAccountRepositoryImpl{DB: DB}
}

func (repo *userAccountRepositoryImpl) CreateUser(newUser entities.UserAccount, rolId uint8) error {
	newUser.Id = uuid.New()
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newUser).Error; err != nil {
			tx.Rollback()
			return err
		}

		rol := entities.UserHasRole{
			UserId: newUser.Id,
			RolId:  rolId,
		}

		if err := tx.Create(&rol).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
