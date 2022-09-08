package custom_error

type TransferToSameWalletError struct {
}

func (e *TransferToSameWalletError) Error() string {
	return "cannot transfer to same wallet"
}
