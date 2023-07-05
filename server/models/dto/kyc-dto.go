package dto

import (
	"time"

	"github.com/google/uuid"
)

type KYCReviewDTOModel struct {
	Id             uuid.UUID
	UserRef        string
	PreRevStatus   string
	CreatedAt      time.Time
	ReviewStatusId uint8
}
