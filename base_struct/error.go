package base_struct

type ValidateError struct {
	Message string
}

func (validateError ValidateError) Error() string {
	return validateError.Message
}
