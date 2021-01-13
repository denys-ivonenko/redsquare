package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type square struct {
	X     int `validate:"required"`
	Y     int `validate:"required"`
	Width int `validate:"required"`
}

//MoveSquare ....
func MoveSquare(c echo.Context) error {
	// get data from request
	var s square
	err := json.NewDecoder(c.Request().Body).Decode(&s)
	if err != nil {
		return err
	}

	// validate data
	err = validate.Struct(s)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// insert data to storage

	return c.NoContent(http.StatusCreated)
}
