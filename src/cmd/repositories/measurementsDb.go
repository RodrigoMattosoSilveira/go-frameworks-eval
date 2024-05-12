package repositories

import (
	"database/sql"
	"fmt"
	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/storage"
)

func CreateMeasurements(Measurements models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()

	result, err := db.Exec("INSERT INTO Measurements (user_id, weight, height, body_fat) VALUES (?, ?, ?, ?)", 
		Measurements.UserId, 
		Measurements.Weight, 
		Measurements.Height, 
		Measurements.BodyFat)
	if err != nil {
		return Measurements, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Measurements, err
	}

	row := db.QueryRow("SELECT * FROM Measurements WHERE Id = ?", id)
    if err := row.Scan(
			&Measurements.Id, 
			&Measurements.UserId, 
			&Measurements.Weight, 
			&Measurements.Height, 
			&Measurements.BodyFat, 
			&Measurements.CreatedAt,
			&Measurements.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return Measurements, fmt.Errorf("id %d: no such Measurements", id)
        }
        return Measurements, fmt.Errorf("id %d: %v", id, err)
    }

	return Measurements, nil
}