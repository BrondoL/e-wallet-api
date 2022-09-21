package custom_error

type InsufficientBallanceError struct {
}

func (e *InsufficientBallanceError) Error() string {
	return "insufficient ballance"
}
