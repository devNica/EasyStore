package impl

import (
	"github.com/devnica/EasyStore/entities"
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
