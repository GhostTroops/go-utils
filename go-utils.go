package go_utils

import (
	"embed"
	"sync"
)

//go:embed config/*
var config embed.FS

func init() {
	Wg = &sync.WaitGroup{}
	DoInit(&config)
}
