package utils

import "fmt"

func GenerateWalletNumber(userID uint) string {
	if userID < 10 {
		return fmt.Sprintf("10000%d", userID)
	} else if userID < 100 {
		return fmt.Sprintf("1000%d", userID)
	} else {
		return fmt.Sprintf("100%d", userID)
	}
}
