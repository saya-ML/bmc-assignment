package api

import (
	"bmc/bmc-assignment/cmd/config"
	log "github.com/sirupsen/logrus"
)

func StartApp() error {
	log.Info("Starting App")
	dependencies := config.BuildDependencies()
	router := Build(dependencies)
	if routerErr := router.Run(":8080"); routerErr != nil {
		log.Error("Error starting router", routerErr)
		return routerErr
	}
	return nil
}
