package impl

import (
	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/devnica/EasyStore/repositories"
	"gorm.io/gorm"
)

type adminCommitRepositoryImpl struct {
	*gorm.DB
}

func NewAdminCommitRepositoryIMpl(DB *gorm.DB) repositories.AdminCommitRepository {
	return &adminCommitRepositoryImpl{DB: DB}
}

func (repo *adminCommitRepositoryImpl) InsertKYCReviewHistory(data dto.KYCReviewDTOModel) error {

	kycReview := entities.KYCReviewRequest{
		Id:             data.Id,
		UserRef:        data.UserRef,
		PreRevStatus:   data.PreRevStatus,
		CreatedAt:      data.CreatedAt,
		ReviewStatusId: data.ReviewStatusId,
	}

	err := repo.DB.Create(&kycReview).Error
	if err != nil {
		return err
	}
	return nil
}
