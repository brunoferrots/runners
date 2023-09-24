package services

import (
	"net/http"
	"p_runners/models"
	"strconv"
	"time"
)

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
	responseErr := validatedRunner(runner)
	if responseErr != nil {
		return nil, responseErr
	}
	return rs.resultsRepositories.CreateRunner(runner)
}

func (rs RunnersService) UpdateRunner(runner *models.Runner) *models.ResponseError {
	responseErr := validateRunnerId(runner.ID)
	if responseErr != nil {
		return responseErr
	}

	responseErr = validatedRunner(runner)
	if responseErr != nil {
		return responseErr
	}

	return rs.runnersRepositories.UpdateRunner(runner)
}

func (rs RunnersService) DeleteRunner(runnerId string) *models.ResponseError {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return responseErr
	}

	return rs.runnersRepositories.DeleteRunner(runnerId)
}

func (rs RunnersService) GetRunner(runnerId string) (*models.Runner, models.ResponseError) {
	responseErr := validateRunnerId(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}

	runner, responseErr := rs.runnersRepositories.GetRunner(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}

	results, responseErr := rs.resultsRepositories.GetAllResults(runnerId)
	if responseErr != nil {
		return nil, responseErr
	}

	runner.Result = results
	return runner

}

func (rs RunnersService) GetRunnersBatch(country string, year string) ([]*models.Runner, *models.ResponseError) {
	if country != "" && year != "" {
		return nil, &models.ResponseError{
			Message: "Only one parameter can be passed",
			Status:  http.StatusBadRequest,
		}
	}

	if country != "" {
		return rs.runnersRepositories.GetRunnersByCountry(country)
	}

	if year != "" {
		intYear, err := strconv.Atoi(year)
		if err != nil {
			return nil, &models.ResponseError{
				Message: "Invalid year",
				Status:  http.StatusBadRequest,
			}
		}
		currentYear := time.Now().Year()
		if intYear < 0 || intYear > currentYear {
			return nil, &models.ResponseError{
				Message: "Invalid year",
				Status:  http.StatusBadRequest,
			}
		}
		return rs.runnersRepositories.GetRunnersByYear(intYear)
	}

	return rs.runnersRepositories.GetAllRunners()
}

func validatedRunner(runner *models.Runner) *models.ResponseError {
	if runner.FirstName == "" {
		return &models.ResponseError{
			Message: "First name is empty",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.LastName == "" {
		return &models.ResponseError{
			Message: "Last name is empty",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.Age <= 16 || runner.Age >= 122 {
		return &models.ResponseError{
			Message: "Age should be between 16 and 122",
			Status:  http.StatusBadRequest,
		}
	}

	if runner.Country == "" {
		return &models.ResponseError{
			Message: "Country is empty",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}

func validateRunnerId(runnerId string) *models.ResponseError {
	if runnerId == "" {
		return &models.ResponseError{
			Message: "Runner id is empty",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
