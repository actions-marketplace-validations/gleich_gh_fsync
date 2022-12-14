package utils

import (
	"os"
	"testing"

	"github.com/gleich/logoru"
)

// Check an error in one line
func CheckErr(msg string, err error) {
	if err != nil {
		logoru.Critical(msg, ";", err)
		os.Exit(1)
	}
}

// Check for a error in one line
func CheckTestingErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
