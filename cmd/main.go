package main

import (
	"github.com/api-voting/config"
	"github.com/api-voting/internal/app/delivery"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}

func main() {
	delivery.Server().Run()
}
