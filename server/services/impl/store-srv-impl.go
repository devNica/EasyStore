package impl

import (
	"context"

	"github.com/devnica/EasyStore/exceptions"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/devnica/EasyStore/models/request"
	"github.com/devnica/EasyStore/models/response"
	"github.com/devnica/EasyStore/repositories"
	"github.com/devnica/EasyStore/services"
	"github.com/google/uuid"
	"github.com/mmcloughlin/geohash"
)

type storeServiceImpl struct {
	repositories.StoreRepository
}

func NewStoreServiceImpl(repo *repositories.StoreRepository) services.StoreService {
	return &storeServiceImpl{
		StoreRepository: *repo,
	}
}

func (srv *storeServiceImpl) RegisterStore(ctx context.Context, newStore request.StoreRequestModel, userId string) {

	ownerId, err := uuid.Parse(userId)

	exceptions.PanicLogging(err)

	storeDTO := dto.StoreRegisterDTOModel{
		Id:        uuid.New(),
		StoreName: newStore.StoreName,
		Address:   newStore.Address,
		NIT:       newStore.NIT,
		GeoHash:   geohash.Encode(newStore.Latitude, newStore.Longitude),
		OwnerId:   ownerId,
	}

	err = srv.StoreRepository.CreateStore(storeDTO)
	exceptions.PanicLogging(err)
}

func (srv *storeServiceImpl) GetStoresByOwnerId(ctx context.Context, ownerId string) []response.StoreResponseModel {

	result, err := srv.StoreRepository.FetchStoresByOwnerId(ownerId)
	exceptions.PanicLogging(err)

	var stores []response.StoreResponseModel
	for _, store := range result {
		lat, lng := geohash.Decode(store.GeoHash)

		stores = append(stores, response.StoreResponseModel{
			StoreId:   store.StoreId,
			StoreName: store.StoreName,
			Address:   store.Address,
			NIT:       store.NIT,
			IsActive:  store.IsActive,
			CreatedAt: store.CreatedAt,
			Latitude:  lat,
			Longitude: lng,
		})
	}

	return stores

}
