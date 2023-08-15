package utils

import (
	"strings"
)

func ExtractDuplicateKeyError(errorMessage string) string {
	duplicateKey := strings.Split(errorMessage, "dup key: { ")
	if len(duplicateKey) != 2 {
		return ""
	}
	return strings.Split(duplicateKey[1], ": ")[0]
}
