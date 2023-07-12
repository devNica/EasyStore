package commons

var dict = make(map[string]interface{})

var roles = map[string]uint8{
	"admins":    1,
	"owners":    2,
	"customers": 3,
	"operators": 4,
}

var accountStatus = map[string]uint8{
	"unverifiableIdentity": 1,
	"awaitingReview":       2,
	"approved":             3,
	"rejected":             4,
	"locked":               5,
}

var reviewStatus = map[string]uint8{
	"pending assigment":  1,
	"review in progress": 2,
	"confirming changes": 3,
	"review finished":    4,
	"reassigned review":  5,
}

func GetRolesFromDictionary() interface{} {
	dict["roles"] = roles
	return dict["roles"]
}

func GetAccountStatusFromDictionary() interface{} {
	dict["status"] = accountStatus
	return dict["status"]
}

func GetReviewStatusFromDictionary() interface{} {
	dict["reviewStatus"] = reviewStatus
	return dict["reviewStatus"]
}

func GetKeyId(key string, dictionary interface{}) uint8 {
	i := dictionary.(map[string]uint8)
	return i[key]
}
