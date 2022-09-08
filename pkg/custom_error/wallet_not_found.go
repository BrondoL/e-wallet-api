package custom_error

type WalletNotFoundError struct {
}

func (e *WalletNotFoundError) Error() string {
	return "wallet not found"
}
