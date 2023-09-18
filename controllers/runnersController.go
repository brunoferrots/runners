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
		log.Println("Error while reading create runners request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		log.Println("Erro while unmarshaling update runners request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := rh.runnersService.CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rh RunnersController) UpdateRunner(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading update runners request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var runner models.Runner
	err = json.Unmarshal(body, &runner)
	if err != nil {
		log.Println("Erro while unmarshaling update runners request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, reponseErr := rh.runnersService.UpdateRunner(&runner)
	if reponseErr != nil {
		ctx.AbortWithStatusJSON(reponseErr.Status, reponseErr)
		return
	}

	ctx.JSON(http.StatusNoContent, response)
}

func (rh RunnersController) DeleteRunner(ctx *gin.Context) {
	runnerId := ctx.Param("id")
	responseErr := rh.runnersService.DeleteRunner(runnerId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (rh RunnersController) GetRunner(ctx *gin.Context) {

}

func (rh RunnersController) GetRunnersBatch(ctx *gin.Context) {

}
