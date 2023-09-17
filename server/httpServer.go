package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	runnersController *controllers.RunnersController
	resultsController *controllers.ResultController
}

func InitHttpServer(config *viper.Viper, dbHandle *sql.DB) HttpServer {
	runnerRepository := repositories.NewRunnersRepository(dbHandle)
	resultRepository := repositories.NewResultRepository(dbHandle)

	runnersService := services.NewRunnersService(runnerRepository, resultRepository)
	resultsService := services.NewResultsService(resultRepository, runnerRepository)

	runnersController := controllers.NewRunnersController(runnersService)
	resultsController := controllers.NewResultsController(resultsService)

	router := gin.Default()
	router.POST("/runner", runnersController.CreateRunner)
	router.PUT("/runner", runnersController.UpdateRunner)
	router.DELETE("/runner/:id", runnersController.DeleteRunner)
	router.GET("/runner/:id", runnersController.GetRunner)
	router.GET("/runner", runnersController.GetRunnersBatch)

	router.POST("/result", resultsController.CreateResult)
	router.DELETE("/result/:id", resultsController.DeleteResult)

	return HttpServer{
		config:            config,
		router:            router,
		runnersController: runnersController,
		resultsController: resultsController,
	}
}

func (hs HttpServer) Start() {

}
