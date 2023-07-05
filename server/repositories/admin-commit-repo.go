package repositories

import "github.com/devnica/EasyStore/models/dto"

type AdminCommitRepository interface {
	InsertKYCReviewHistory(data dto.KYCReviewDTOModel) error
}
