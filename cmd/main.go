package main

import (
	"log"

	"github.com/skvdmt/skvdmt-e2e-tests/internal"
)

// Точка входа в приложение.
func main() {
	a, err := internal.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
