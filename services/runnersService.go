package services

import "p_runners/models"

type RunnersService struct {
	runnersRepositories *repositories.RunnersRepository
	resultsRepositories *repositories.ResultsRepository
}

func NewRunnersService(runnersRepositories *repositories.RunnersRepository, resultsRepositories *repositories.ResultsRepository) *RunnersService {
	return &RunnersService{
		runnersRepositories: runnersRepositories,
		resultsRepositories: resultsRepositories,
	}
}

func (rs RunnersService) CreateRunner(runner *models.Runner) (*models.Runner, *models.ResponseError) {

}

func (rs RunnersService) UpdateRunner(runner *models.Runner) *models.ResponseError {

}

func (rs RunnersService) DeleteRunner(runnerId string) *models.ResponseError {

}

func (rs RunnersService) GetRunner(runnerId string) (*models.Runner, models.ResponseError) {

}

func (rs RunnersService) GetRunnersBatch(country string, year int) ([]*models.Runner, models.ResponseError) {

}
