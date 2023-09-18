package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"p_runners/models"

	"github.com/gin-gonic/gin"
)

type RunnersController struct {
	runnersService *runnersService
}

func NewRunnersController(runnersService *service.runnersService) *RunnersController {
	return &RunnersController{
		runnersService: runnersService,
	}
}

func (rh RunnersController) CreateRunner(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading runners request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {

	}
}

func (rh RunnersController) UpdateRunner(ctx *gin.Context) {

}

func (rh RunnersController) DeleteRunner(ctx *gin.Context) {

}

func (rh RunnersController) GetRunner(ctx *gin.Context) {

}

func (rh RunnersController) GetRunnersBatch(ctx *gin.Context) {

}
