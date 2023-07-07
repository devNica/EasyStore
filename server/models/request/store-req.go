package request

type KeyComposedUserStoreModel struct {
	OwnerId string `json:"ownerId"`
	StoreId string `json:"storeId"`
}

type StoreRequestModel struct {
	StoreName string  `json:"storeName" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	NIT       string  `json:"nit" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type UpdateStoreRequestModel struct {
	StoreName string  `json:"storeName" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	NIT       string  `json:"nit" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}
