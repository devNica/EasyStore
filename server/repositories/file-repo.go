package repositories

import (
	"github.com/devnica/EasyStore/models/dto"
	"github.com/google/uuid"
)

type FileRepository interface {
	InsertStoreAsset(asset dto.InsertFileDTOModel, storeId uuid.UUID) error
}
