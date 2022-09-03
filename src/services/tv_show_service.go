package services

import (
	"errors"

	"github.com/Israel-Ferreira/binge-watchers/src/dto"
	"github.com/Israel-Ferreira/binge-watchers/src/models"
	"github.com/Israel-Ferreira/binge-watchers/src/repositories"
)

func ValidationError(msg string) error {
	return errors.New(msg)
}

type TvShowService struct {
	tvShowRepo repositories.TvShowRepository
}

func (tvs TvShowService) FindAll() ([]models.TvShow, error) {
	tvShows, err := tvs.tvShowRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return tvShows, nil
}

func (tvs TvShowService) FindById(id uint64) (*models.TvShow, error) {
	tvShow, err := tvs.tvShowRepo.FindById(id)

	if err != nil {
		return &models.TvShow{}, err
	}

	return tvShow, nil
}

func (tvs TvShowService) DeleteById(id uint64) error {
	return tvs.tvShowRepo.DeleteById(id)
}

func (tvs TvShowService) Update(id uint64, dto models.TvShow) error {
	return nil
}

func (tvs TvShowService) Create(dto dto.SerieDTO) (models.TvShow, error) {

	if dto.Title == "" {
		return models.TvShow{}, ValidationError("O titulo não pode estar em branco")
	}

	if dto.Category == "" {
		return models.TvShow{}, ValidationError("A categoria da serie não pode estar em branco")
	}

	tvShow := models.TvShow{
		Title:      dto.Title,
		Category:   dto.Category,
		LaunchYear: dto.LaunchYear,
		IsFinished: dto.IsFinished,
	}

	createdTvShow, err := tvs.tvShowRepo.Create(tvShow)

	if err != nil {
		return models.TvShow{}, err
	}

	return *createdTvShow, nil
}

func NewTvShowService(repo repositories.TvShowRepository) TvShowService {
	return TvShowService{tvShowRepo: repo}
}
