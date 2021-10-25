package tests_test

import (
	"testing"

	"gitlab.com/Dank-del/SibylAPI-Go/sibyl/core/utils/logging"
)

// test to create a token using the owner token with http
// headers. The token should be created successfully.
func TestCreateToken01(t *testing.T) {
	f := logging.LoadLogger()
	t.Cleanup(func() {
		if f != nil {
			f()
		}
	})

	// run the app in anoter goroutine
	go runApp()
}

// test to create a token using the owner token with http
// url encoded variables. The token should be created successfully.
func TestCreateToken02(t *testing.T) {
}
func TestCreateToken03(t *testing.T) {
}
func TestCreateToken04(t *testing.T) {
}
func TestCreateToken05(t *testing.T) {
}
func TestCreateToken06(t *testing.T) {
}
func TestCreateToken07(t *testing.T) {
}
