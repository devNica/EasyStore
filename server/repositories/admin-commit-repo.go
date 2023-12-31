package repositories

import (
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
)

type AdminCommitRepository interface {
	InsertKYCReviewHistory(data dto.KYCReviewDTOModel) error
	FetchAllKYCReview(reviewStatusId uint8) ([]dao.KYCReviewDAOModel, error)
}
