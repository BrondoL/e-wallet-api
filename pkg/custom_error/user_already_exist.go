package custom_error

type UserAlreadyExistsError struct {
}

func (e *UserAlreadyExistsError) Error() string {
	return "user already exists"
}
