package utils

import "log"

import (
	"github.com/hasirciogli/multicloud-csi-driver/node"
)

// LogError logs an error message
func LogError(err error) {
    if err != nil {
        log.Printf("Error: %v", err)
    }
}

// Other utility functions... 