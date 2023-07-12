package impl

import (
	"errors"
	"fmt"

	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/models/dao"
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

func (repo *adminCommitRepositoryImpl) FetchAllKYCReview(reviewStatusID uint8) ([]dao.KYCReviewDAOModel, error) {
	var kyc []dao.KYCReviewDAOModel

	result := repo.DB.Table("kyc_review_request").
		Select(`
		kyc_review_request.user_ref,
		kyc_review_request.admin_ref,
		kyc_review_request.pre_rev_status,
		kyc_review_request.post_rev_status,
		kyc_review_request.is_assigned,
		kyc_review_request.created_at,
		kyc_review_request.updated_at,
		review_status.id as "ReviewStatusId"
	`).
		Joins("inner join review_status on review_status.id = kyc_review_request.review_status_id").
		Where("review_status.id = ?", reviewStatusID).Scan(&kyc)

	fmt.Println(kyc)

	if result.RowsAffected == 0 {
		return []dao.KYCReviewDAOModel{}, errors.New("No kyc list was retrieved for review")
	}

	return kyc, nil

}
