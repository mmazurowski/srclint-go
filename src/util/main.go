package util

import "os"

func CheckExit(err error, status int) {
	if err != nil {
		os.Exit(status)
	}
}
