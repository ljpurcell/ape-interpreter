package utils

import (
	"log"
)

func Check(err error, msg string) {
    if err != nil {
        log.Fatalf(msg + ": %v", err)
    }
}

func IsLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func IsDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
