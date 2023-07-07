package impl

import (
	"github.com/devnica/EasyStore/entities"
	"github.com/devnica/EasyStore/models/dto"
	"github.com/devnica/EasyStore/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type fileRepositoryImpl struct {
	*gorm.DB
}

func NewFileRepositoryImpl(DB *gorm.DB) repositories.FileRepository {
	return &fileRepositoryImpl{DB: DB}
}

func (repo *fileRepositoryImpl) InsertStoreAsset(asset dto.InsertFileDTOModel, storeId uuid.UUID) error {
	newAsset := entities.File{
		Filename:  asset.Filename,
		Filetype:  asset.Filetype,
		Filesize:  asset.Filesize,
		Binary:    asset.Binary,
		CreatedAt: asset.CreatedAt,
	}

	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newAsset).Error; err != nil {
			tx.Rollback()
			return err
		}

		rol := entities.StoreHasAssets{
			StoreId: storeId,
			FileId:  newAsset.Filename,
		}

		if err := tx.Create(&rol).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
