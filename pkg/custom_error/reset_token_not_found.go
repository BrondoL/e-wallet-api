package custom_error

type ResetTokenNotFound struct {
}

func (e *ResetTokenNotFound) Error() string {
	return "invalid reset token"
}
