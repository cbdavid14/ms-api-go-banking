package main

import (
	"github.com/cbdavid14/ms-api-go-banking/app"
	"github.com/cbdavid14/ms-api-go-banking/logger"
)

func main() {

	logger.Info("Starting the application...")
	app.Start()
}
