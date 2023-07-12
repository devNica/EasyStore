package response

import (
	"time"

	"github.com/google/uuid"
)

type KYCReviewResponseModel struct {
	ReviewId      uuid.UUID `json:"reviewId"`
	UserRef       string    `json:"userRef"`
	AdminRef      string    `json:"adminRef"`
	PreRevStatus  string    `json:"preRevStatus"`
	PostRevStatus string    `json:"postRevStatus"`
	IsAsigned     bool      `json:"isAsigned"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
