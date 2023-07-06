package impl

import (
	"context"

	"github.com/devnica/EasyStore/commons"
	"github.com/devnica/EasyStore/commons/security"
	"github.com/devnica/EasyStore/commons/utils"
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
	repositories.UserAccountRepository
	repositories.StoreRepository
}

func NewStoreServiceImpl(storeRepo *repositories.StoreRepository, accRepo *repositories.UserAccountRepository) services.StoreService {
	return &storeServiceImpl{
		StoreRepository:       *storeRepo,
		UserAccountRepository: *accRepo,
	}
}

func (srv *storeServiceImpl) RegisterStore(
	ctx context.Context,
	newStore request.StoreRequestModel,
	userId string) response.StoreRegisterResponseModel {

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

	/*
		Se recuperan los roles posteriores a la creacion de la tienda
		y luego se recorren para determinar que si el usuario ya tiene el rol de "owners"
		no tenga que actualizar el token devuelto por el controlador,
		pero en caso de que no lo posea se prosigue con la logica de actualizacion de roles
		y generacion de un nuevo token
	*/

	prevRoles, err := srv.UserAccountRepository.FetchRolesByUserId(ownerId.String())
	exceptions.PanicLogging(err)

	for _, role := range prevRoles {
		if role.Role == "owners" {
			return response.StoreRegisterResponseModel{
				Token:          "",
				TokenIsUpdated: false,
			}
		}
	}

	/*
		una vez se realizar el registro de los datos de la tienda
		se procede a actualizar los roles del usuario en la base de datos
		se obtiene del diccionario de datos el ID del rol "owners"
		se invoca al servicio de inserccion de roles a la cuenta de
		usuario, pasandole tanto el id del rol, como el del usuario que
		se va afectar
	*/

	rolesMap := commons.GetRolesFromDictionary()
	roleId := commons.GetKeyId("owners", rolesMap)

	err = srv.UserAccountRepository.InsertRoleToUserAccount(ownerId, roleId)
	exceptions.PanicLogging(err)

	/*
		Terminada la actualizacion de los roles en la base de datos
		se realiza una consulta para recuperar estos roles y con ellos
		procesar un nuevo token que le permita al usuario
		authenticarse en aquellas rutas donde se necesite el rol "owners"
		al que recientemente se acaba de promocionar el usuario cliente
		dicho token se retorna en la respuesta del controlador
	*/

	roles, err := srv.UserAccountRepository.FetchRolesByUserId(ownerId.String())
	exceptions.PanicLogging(err)

	newRolesMap := utils.ConvertRolesToMaps(roles)

	updateToken := security.GenerateToken(ownerId.String(), newRolesMap)

	return response.StoreRegisterResponseModel{
		Token:          updateToken,
		TokenIsUpdated: true,
	}

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
