package main

import (
	"github.com/dharmatin/sim-k-api/internal/config"
	"github.com/dharmatin/sim-k-api/internal/db"
)

func main() {
	cfg, _ := config.Load("./cmd/config.local.yaml")
	db := db.Setup(cfg.Database.Type, cfg.Database.PSN)
	db.AutoMigrate()
}
