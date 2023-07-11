package impl

import (
	"errors"

	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
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

func (repo *userAccountRepositoryImpl) CreateUser(data dto.UserRegisterDTOModel, roleId uint8) error {

	newUser := entities.UserAccount{
		Id:            data.Id,
		Email:         data.Email,
		Password:      data.Password,
		CreatedAt:     data.CreatedAt,
		TwoFactorAuth: data.TwoFactorAuth,
		StatusId:      data.StatusId,
	}

	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newUser).Error; err != nil {
			tx.Rollback()
			return err
		}

		rol := entities.UserHasRoles{
			UserId: newUser.Id,
			RoleId: roleId,
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

	result := repo.DB.Table("roles").
		Select(`
		roles.id,
		roles.role
	`).
		Joins("inner join user_has_roles on user_has_roles.role_id = roles.id").
		Joins("inner join user_account on user_account.id = user_has_roles.user_id").
		Where("user_account.id = ?", userId).Scan(&RolesResult)

	if result.RowsAffected == 0 {
		return []dao.RolDAOModel{}, errors.New("User has no roles assigned")
	}

	return RolesResult, nil

}

func (repo *userAccountRepositoryImpl) InsertPersonalInfo(personalInfo dto.PersonalInfoDTOModel, userId string) error {

	result := repo.DB.Model(&entities.UserAccount{}).Where("id=?", userId).Updates(personalInfo)

	if result.RowsAffected == 0 {
		return errors.New("Personal info update failure")
	}

	return nil
}

func (repo *userAccountRepositoryImpl) InsertRoleToUserAccount(userId uuid.UUID, roleId uint8) error {

	userRoles := entities.UserHasRoles{
		UserId:   userId,
		RoleId:   roleId,
		IsActive: true,
	}

	result := repo.DB.Create(&userRoles)

	if result.RowsAffected == 0 {
		return errors.New("could not add new role to user account")
	}

	return nil
}

func (repo *userAccountRepositoryImpl) FetchStatusAccountByUserId(userId uuid.UUID) (dao.AccounStatusDAOModel, error) {

	var statusResult dao.AccounStatusDAOModel

	result := repo.DB.Table("account_status").
		Select(`
		account_status.status
	`).
		Joins("inner join user_account on user_account.status_id = account_status.id").
		Where("user_account.id = ?", userId).Scan(&statusResult)

	if result.RowsAffected == 0 {
		return dao.AccounStatusDAOModel{}, errors.New("User not found")
	}

	return statusResult, nil
}
