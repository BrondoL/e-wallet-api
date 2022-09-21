package custom_error

type UserNotFoundError struct {
}

func (e *UserNotFoundError) Error() string {
	return "user not found"
}
