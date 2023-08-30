package api

import (
	"bmc/bmc-assignment/cmd/config"
	"bmc/bmc-assignment/cmd/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Build(dependencies config.Dependencies) *gin.Engine {
	router := gin.Default()
	mapRoutes(router, dependencies)
	return router
}

func mapRoutes(r *gin.Engine, dependencies config.Dependencies) {
	//swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	passenger := handlers.NewPassenger(dependencies.PassengerService)
	histogram := handlers.NewHistogram(dependencies.HistogramService)
	r.GET("/passenger/:key", passenger.Get)
	r.GET("/passenger", passenger.All)
	r.GET("/histogram", histogram.HistPlot)

}
