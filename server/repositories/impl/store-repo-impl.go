package impl

import (
	"errors"

	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/models/dao"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/devnica/EasyStore/repositories"
	"gorm.io/gorm"
)

type storeRepositoryImpl struct {
	*gorm.DB
}

func NewStoreRepositoryImpl(DB *gorm.DB) repositories.StoreRepository {
	return &storeRepositoryImpl{DB: DB}
}

func (repo *storeRepositoryImpl) CreateStore(data dto.StoreRegisterDTOModel) error {
	store := entities.Store{
		Id:        data.Id,
		StoreName: data.StoreName,
		Address:   data.Address,
		NIT:       data.NIT,
		GeoHash:   data.GeoHash,
		OwnerId:   data.OwnerId,
		CreatedAt: data.CreatedAt,
	}

	err := repo.DB.Create(&store).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *storeRepositoryImpl) FetchStoresByOwnerId(ownerId string) ([]dao.StoreDAOModel, error) {
	var storeResult []entities.Store

	result := repo.DB.Where("owner_id = ?", ownerId).Find(&storeResult)

	if result.RowsAffected == 0 {
		return []dao.StoreDAOModel{}, nil
	}

	var stores []dao.StoreDAOModel
	for _, store := range storeResult {
		stores = append(stores, dao.StoreDAOModel{
			StoreId:   store.Id,
			StoreName: store.StoreName,
			Address:   store.Address,
			NIT:       store.NIT,
			GeoHash:   store.GeoHash,
			IsActive:  store.IsActive,
			CreatedAt: store.CreatedAt,
			OwnerId:   store.OwnerId,
		})
	}

	return stores, nil
}

func (repo *storeRepositoryImpl) UpdateStoreByStoreId(relation dto.UserRelationShipWithStoreDTO, storeInfo dto.UpdateStoreDTOModel) error {

	result := repo.DB.Model(&entities.Store{}).Where("id=? and owner_id=?", relation.StoreId, relation.OwnerId).Updates(storeInfo)

	if result.RowsAffected == 0 {
		return errors.New("Store info update failure")
	}

	return nil
}
