package handlers

import (
	"net/http"
	"strconv"

	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/repositories"

	"github.com/labstack/echo/v4"
)

func CreateMeasurementsHandler(c echo.Context) error {
	measurements := models.Measurements{}
	c.Bind(&measurements)
	neNMeasurement, err := repositories.CreateMeasurementsDb(measurements)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, neNMeasurement)
}

func UpdateMeasurementsHandler(c echo.Context) error {
	id := c.Param("id")
  
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
	  return c.JSON(http.StatusInternalServerError, err.Error())
	}
  
	measurements := models.Measurements{}
	c.Bind(&measurements)
	updatedMeasurements, err := repositories.UpdateMeasurementsDb(measurements, idInt)
	if err != nil {
	  return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedMeasurements)
  }
  
//   TODO add support for READ including filters, DELETE
//	TODO Add parameter validation logic
// 	TODO explore middleware
//  TODO Learn more about routing