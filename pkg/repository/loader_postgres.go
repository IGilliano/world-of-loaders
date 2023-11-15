package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"worldOfLoaders/pkg/models"
)

type LoaderPostgres struct {
	db *sqlx.DB
}

func NewLoader(db *sqlx.DB) *LoaderPostgres {
	return &LoaderPostgres{db: db}
}

func (l *LoaderPostgres) GetLoaderFromDB(ID int) (models.Loader, error) {
	var loader models.Loader
	query := fmt.Sprintf(`SELECT * FROM loaders WHERE p_id = $1`)
	err := l.db.Get(&loader, query, ID)

	return loader, err
}

func (l *LoaderPostgres) GetRegisteredLoaders() ([]*models.Loader, error) {
	var loaders []*models.Loader
	row := fmt.Sprintf("SELECT * FROM loaders")
	rows, err := l.db.Query(row)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var loader models.Loader
		err = rows.Scan(&loader.ID, &loader.Capacity, &loader.IsDrinking, &loader.Fatigue, &loader.Salary)
		if err != nil {
			return nil, err
		}
		loaders = append(loaders, &loader)
	}
	return loaders, nil
}

func (l *LoaderPostgres) GetLoadersFromDB(loadersID []int) ([]*models.Loader, error) {
	loaders := make([]*models.Loader, 0)
	for _, id := range loadersID {
		var loader models.Loader
		query := fmt.Sprintf(`SELECT * FROM loaders WHERE p_id = $1`)
		err := l.db.Get(&loader, query, id)
		if err != nil {
			return nil, err
		}
		loaders = append(loaders, &loader)
	}
	return loaders, nil
}

func (l *LoaderPostgres) UpdateLoaders(loaders []*models.Loader) error {
	for _, loader := range loaders {
		_, err := l.db.Query("UPDATE loaders SET fatigue = $1 WHERE p_id = $2", loader.Fatigue, loader.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
