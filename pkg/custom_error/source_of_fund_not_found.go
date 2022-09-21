package custom_error

type SourceOfFundNotFoundError struct {
}

func (e *SourceOfFundNotFoundError) Error() string {
	return "source of fund not found"
}
