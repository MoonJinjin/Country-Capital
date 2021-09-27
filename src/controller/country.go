package controller

import (
	"net/http"

	"github.com/MoonJinjin/Country-Capital/model"
	"github.com/MoonJinjin/Country-Capital/storage"
	"github.com/labstack/echo/v4"
)

func GetCountry(c echo.Context) error {
	country, _ := GetRepoCountry()
	return c.JSON(http.StatusOK, country)
}

func GetRepoStudents() ([]model.Country, error) {
	db := storage.GetDBInstance()
	country := []model.Country{}

	if err := db.Find(&country).Error; err != nil {
		return nil, err
	}

	return country, nil
}
