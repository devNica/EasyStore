package impl

import (
	"errors"

	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/devnica/EasyStore/repositories"
	"gorm.io/gorm"
)

type userAccountRepositoryImpl struct {
	*gorm.DB
}

func NewUserAccountRepositoryImpl(DB *gorm.DB) repositories.UserAccountRepository {
	return &userAccountRepositoryImpl{DB: DB}
}

func (repo *userAccountRepositoryImpl) CreateUser(newUser dto.UserRegisterDTOModel, rolId uint8) error {
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

func (repo *userAccountRepositoryImpl) FindUserByEmail(email string) (dao.FindUserDAOModel, error) {

	var user entities.UserAccount

	result := repo.DB.Select("id, email, password").Where("email = ?", email).First(&user)

	if result.RowsAffected == 0 {
		return dao.FindUserDAOModel{}, errors.New("User not found")
	}

	return dao.FindUserDAOModel{
		Id:       user.Id,
		Email:    user.Email,
		Password: user.Password,
	}, nil

}

func (repo *userAccountRepositoryImpl) FetchRolesByUserId(userId string) ([]dao.RolDAOModel, error) {

	var RolesResult []dao.RolDAOModel

	result := repo.DB.Table("rol").
		Select(`
		rol.id,
		rol.rol
	`).
		Joins("inner join user_has_role on user_has_role.rol_id = rol.id").
		Joins("inner join user_account on user_account.id = user_has_role.user_id").
		Where("user_account.id = ?", userId).Scan(&RolesResult)

	if result.RowsAffected == 0 {
		return []dao.RolDAOModel{}, errors.New("User has no roles assigned")
	}

	return RolesResult, nil

}
