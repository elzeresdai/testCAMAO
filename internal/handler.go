package internal

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

var ToDoS []*ToDo

func SaveNewRecord(ctx echo.Context) error {

	err, rec := GetRecord(ctx)

	if err != nil {
		return err
	}

	ToDoS = append(ToDoS, rec)

	return nil
}

func UpdateRecord(ctx echo.Context) error {

	err, rec := GetRecord(ctx)

	if err != nil {
		return err
	}

	for _, val := range ToDoS {
		if val.Title == rec.Title {
			val.Status = rec.Status
		}
	}

	return nil
}

func GetRecord(ctx echo.Context) (error, *ToDo) {
	record := ToDo{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&record)
	defer ctx.Request().Body.Close()

	if err != nil {
		return err, nil
	}

	return nil, &record
}
