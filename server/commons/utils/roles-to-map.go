package utils

import "github.com/devnica/EasyStore/entities"

func ConvertRolesToMaps(roles []entities.Rol) []map[string]interface{} {
	var rolesMaps []map[string]interface{}

	for _, rol := range roles {
		rolMap := make(map[string]interface{})
		rolMap["rolId"] = rol.Id
		rolMap["rol"] = rol.Rol

		rolesMaps = append(rolesMaps, rolMap)
	}

	return rolesMaps
}
