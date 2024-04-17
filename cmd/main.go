package main

import (
	"github.com/new-task/api"
	"github.com/new-task/config"
	"github.com/new-task/pkg/db"
	"github.com/new-task/pkg/logger"
	"github.com/new-task/storage"
)

func main() {
	cfg := config.Load()
	log := logger.New("debug", "newtask")

	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Info("errors", logger.Error(err))
	}
	strg := storage.NewStoragePg(connDb)

	apiServer := api.New(&api.Options{
		Cfg:     cfg,
		Storage: strg,
		Log:     log,
	})

	if err := apiServer.Run(":" + cfg.HttpPort); err != nil {
		log.Fatal("failed to run server: %v", logger.Error(err))
	}

}
