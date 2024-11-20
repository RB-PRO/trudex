package main

import (
	"trudex/trud_contact/internal/entrypoint"
)

func main() {
	service := entrypoint.NewService()
	service.Run()
}
