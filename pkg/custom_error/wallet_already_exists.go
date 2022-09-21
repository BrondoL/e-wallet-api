package custom_error

type WalletAlreadyExistsError struct {
}

func (e *WalletAlreadyExistsError) Error() string {
	return "wallet already exists"
}
