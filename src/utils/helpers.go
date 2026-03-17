package utils

import (
    "log"
)

func LogError(message string) {
    log.Printf("ERROR: %s", message)
}