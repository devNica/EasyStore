package services

import (
	"context"

	"github.com/devnica/EasyStore/models/request"
)

type StoreService interface {
	RegisterStore(ctx context.Context, newStore request.StoreRequestModel, userId string)
}
