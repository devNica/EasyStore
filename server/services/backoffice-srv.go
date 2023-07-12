package services

import (
	"context"

	"github.com/devnica/EasyStore/models/response"
)

type BackofficeService interface {
	GetKYCReview(ctx context.Context, reviewStatus string) []response.KYCReviewResponseModel
}
