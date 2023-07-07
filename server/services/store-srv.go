package services

import (
	"context"

	"github.com/devnica/EasyStore/models/request"
	"github.com/devnica/EasyStore/models/response"
)

type StoreService interface {
	RegisterStore(ctx context.Context, newStore request.StoreRequestModel, asset request.FileRequestModel, userId string) response.StoreRegisterResponseModel
	GetStoresByOwnerId(ctx context.Context, ownerId string) []response.StoreResponseModel
	UpdateStoreInfoByStoreId(ctx context.Context, relation request.KeyComposedUserStoreModel, data request.UpdateStoreRequestModel)
}
