package exceptions

type UnAuthorizedError struct {
	Message string
}

func (unAuthorizedError UnAuthorizedError) Error() string {
	return unAuthorizedError.Message
}
