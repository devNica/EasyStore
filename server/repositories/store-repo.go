package repositories

import (
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
)

type StoreRepository interface {
	CreateStore(data dto.StoreRegisterDTOModel) error
	FetchStoresByOwnerId(ownerId string) ([]dao.StoreDAOModel, error)
	UpdateStoreByStoreId(relation dto.KeyComposeUserStoreDTOModel, storeInfo dto.UpdateStoreDTOModel) error
}
