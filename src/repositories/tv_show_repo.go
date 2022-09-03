package repositories

import (
	"github.com/Israel-Ferreira/binge-watchers/src/database"
	"github.com/Israel-Ferreira/binge-watchers/src/models"
)

type TvShowRepository interface {
	FindById(uint64) (*models.TvShow, error)
	FindAll() ([]models.TvShow, error)
	DeleteById(uint64) error
	Create(models.TvShow) (*models.TvShow, error)
	Update(uint64, models.TvShow) error
}

type TvShowRepo struct {
}

func (tvsr TvShowRepo) FindById(id uint64) (*models.TvShow, error) {
	db, err := database.OpenConnection()

	if err != nil {
		return &models.TvShow{}, err
	}

	defer db.Close()

	query, err := db.Query(
		`select tvs.* from tv_shows tvs where tvs.id = $1`,
		id,
	)

	if err != nil {
		return &models.TvShow{}, err
	}

	defer query.Close()

	var serie models.TvShow

	if query.Next() {
		if err := query.Scan(&serie.ID, &serie.Title, &serie.Category, &serie.IsFinished, &serie.LaunchYear); err != nil {
			return &models.TvShow{}, err
		}
	}

	return &serie, nil
}

func (tvsr TvShowRepo) FindAll() ([]models.TvShow, error) {
	db, err := database.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	query, err := db.Query("select tvs.* from tv_shows tvs")

	if err != nil {
		return nil, err
	}

	defer query.Close()

	var series []models.TvShow

	for query.Next() {
		var serie models.TvShow

		if err := query.Scan(&serie.ID, &serie.Title, &serie.Category, &serie.IsFinished, &serie.LaunchYear); err != nil {
			return nil, err
		}

		series = append(series, serie)
	}

	return series, nil
}

func (tvsr TvShowRepo) Create(tvShow models.TvShow) (*models.TvShow, error) {
	db, err := database.OpenConnection()

	if err != nil {
		return &models.TvShow{}, err
	}

	defer db.Close()

	stmt, err := db.Prepare(
		`insert into tv_shows (title, category, is_finished, launch_year)
		VALUES ($1, $2, $3, $4)
		`,
	)

	if err != nil {
		return &models.TvShow{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(tvShow.Title, tvShow.Category, tvShow.IsFinished, tvShow.LaunchYear)

	if err != nil {
		return &models.TvShow{}, nil
	}

	return &tvShow, nil
}

func (tvsr TvShowRepo) DeleteById(id uint64) error {
	db, err := database.OpenConnection()

	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("delete from tv_shows where id = $1")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (tvsr TvShowRepo) Update(id uint64, tvShow models.TvShow) error {
	return nil
}

func NewTvShowRepo() TvShowRepository {
	return TvShowRepo{}
}
