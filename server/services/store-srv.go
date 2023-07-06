package services

import (
	"context"

	"github.com/devnica/EasyStore/models/request"
	"github.com/devnica/EasyStore/models/response"
)

type StoreService interface {
	RegisterStore(ctx context.Context, newStore request.StoreRequestModel, userId string)
	GetStoresByOwnerId(ctx context.Context, ownerId string) []response.StoreResponseModel
}
