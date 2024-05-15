package repositories

import (
	"database/sql"
	"fmt"

	"madronetek.com/go-frameworks-eval/cmd/models"
	"madronetek.com/go-frameworks-eval/cmd/storage"
)

func CreateMeasurementsDb(measurements models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()

	// TODO do not insert a measurements record for a user that already has one
	// select exists (select * from measurements where user_id = measurements.UserId);
	// 
	result, err := db.Exec("INSERT INTO Measurements (user_id, weight, height, body_fat) VALUES (?, ?, ?, ?)",
		measurements.UserId,
		measurements.Weight,
		measurements.Height,
		measurements.BodyFat)
	if err != nil {
		return measurements, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return measurements, err
	}

	row := db.QueryRow("SELECT * FROM Measurements WHERE Id = ?", id)
	if err := row.Scan(
		&measurements.Id,
		&measurements.UserId,
		&measurements.Weight,
		&measurements.Height,
		&measurements.BodyFat,
		&measurements.CreatedAt,
		&measurements.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return measurements, fmt.Errorf("id %d: no such Measurements", id)
		}
		return measurements, fmt.Errorf("id %d: %v", id, err)
	}

	return measurements, nil
}

func UpdateMeasurementsDb(measurements models.Measurements, id int64) (models.Measurements, error) {
	db := storage.GetDB()

	// TODO only update a measurements record for a user that already has one
	// select exists (select * from measurements where user_id = measurements.UserId);
	// 
	sqlStatement := `
	UPDATE measurements 
	SET 
	  weight = ?, 
	  height = ?, 
	   body_fat = ?
	WHERE user_id = ?`
	_, err := db.Exec(sqlStatement, measurements.Weight, measurements.Height, measurements.BodyFat, id)
	if err != nil {
	  return measurements, err
	}
  
	row := db.QueryRow("SELECT * FROM measurements WHERE user_id = ?", id)
	if err := row.Scan(&measurements.Id, &measurements.UserId, &measurements.Weight, &measurements.Height, &measurements.BodyFat, &measurements.CreatedAt, &measurements.UpdatedAt); err != nil {
	  if err == sql.ErrNoRows {
		return measurements, fmt.Errorf("user_id %d: no such measurements", measurements.UserId)
	  }
	
	  return measurements, fmt.Errorf("user_id %d: %v", measurements.UserId, err)
	}
  
	return measurements, nil
  }
  