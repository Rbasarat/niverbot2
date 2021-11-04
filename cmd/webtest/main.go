package main

import (
	"github.com/rbasarat/niverobot/internal/mocksource"
	"time"
)

func main() {

	_ = mocksource.NewMockSource(mocksource.Settings{})

	time.Sleep(10 * time.Minute)
}
