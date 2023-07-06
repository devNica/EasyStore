package utils

import (
	"github.com/devnica/EasyStore/models/dao"
)

func ConvertRolesToMaps(roles []dao.RolDAOModel) []map[string]interface{} {
	var rolesMaps []map[string]interface{}

	for _, rol := range roles {
		rolMap := make(map[string]interface{})
		rolMap["rolId"] = rol.Id
		rolMap["role"] = rol.Role

		rolesMaps = append(rolesMaps, rolMap)
	}

	return rolesMaps
}
