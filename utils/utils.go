package utils

import (
	"fmt"
	"os"
)

var Prefix = Blue + "[present-aur] " + Reset

func ClearCache() {
	cacheDir, err := os.UserCacheDir() 
	if err != nil {
		fmt.Println(Prefix + "Error getting user's cache directory: \n", err)
		os.Exit(0)
	}

	clearError := os.RemoveAll(cacheDir + "/present-aur/")
	if clearError != nil {
		fmt.Println(Prefix + "Error clearing user's cache:\n", clearError)
		os.Exit(0)
	}
}
