// main_test.go

package main

import (
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Init()

	code := m.Run()
	os.Exit(code)
}