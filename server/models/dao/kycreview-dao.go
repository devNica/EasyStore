package dao

import (
	"time"

	"github.com/google/uuid"
)

type KYCReviewDAOModel struct {
	ReviewId       uuid.UUID
	UserRef        string
	AdminRef       string
	PreRevStatus   string
	PostRevStatus  string
	IsAsigned      bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ReviewStatusId uint8
}
