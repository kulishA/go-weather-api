package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kulishA/go-weather-api/domain"
)

type CityPostgres struct {
	db *sqlx.DB
}

func NewCityPostgres(db *sqlx.DB) *CityPostgres {
	return &CityPostgres{db: db}
}

func (c *CityPostgres) Get(userId int) ([]domain.City, error) {
	return nil, nil
}

func (c *CityPostgres) Save(userId int, cityId int) (bool, error) {
	return false, nil
}

func (c *CityPostgres) SaveAll(cities []domain.City) error {
	queryInsert := fmt.Sprintf("INSERT INTO %s (city_id, name, region, country, lat, lon, url) values ", citiesTable)
	var insertParams []interface{}
	for i, city := range cities {
		pos := i * 7
		queryInsert += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d),", pos+1, pos+2, pos+3, pos+4, pos+5, pos+6, pos+7)
		insertParams = append(insertParams, city.CityId, city.Name, city.Region, city.Country, city.Lat, city.Lon, city.Url)
	}
	queryInsert = queryInsert[:len(queryInsert)-1]
	queryInsert += " ON CONFLICT (city_id) DO NOTHING"

	_, err := c.db.Exec(queryInsert, insertParams...)
	if err != nil {
		return err
	}
	return nil
}
