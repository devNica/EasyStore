package repositories

import "github.com/devnica/EasyStore/models/dto"

type StoreRepository interface {
	CreateStore(data dto.StoreRegisterDTOModel) error
}
