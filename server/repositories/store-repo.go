package repositories

import (
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
)

type StoreRepository interface {
	CreateStore(data dto.StoreRegisterDTOModel) error
	FetchStoresByOwnerId(ownerId string) ([]dao.StoreDAOModel, error)
	UpdateStoreByStoreId(storeId string, storeInfo dto.UpdateStoreDTOModel) error
}
