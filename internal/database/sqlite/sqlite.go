package sqlite

import (
	"bmc/bmc-assignment/internal/database"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type dbSQLite struct {
	filename string
}

func (d *dbSQLite) Get(key string) (database.Passenger, error) {
	db, err := gorm.Open(sqlite.Open(d.filename), &gorm.Config{})
	if err != nil {
		return database.Passenger{}, err
	}
	passenger := &database.Passenger{}
	err = db.First(&passenger, "passenger_id = ?", key).Error
	if err != nil {
		return database.Passenger{}, err
	}
	return *passenger, nil
}

func (d *dbSQLite) All() ([]database.Passenger, error) {
	passengers := make([]database.Passenger, 0)
	db, err := gorm.Open(sqlite.Open(d.filename), &gorm.Config{})
	if err != nil {
		return passengers, err
	}
	err = db.Model(&database.Passenger{}).
		Find(&passengers).
		Error
	return passengers, err
}

func New(filename string) database.IDatabase {
	return &dbSQLite{
		filename: filename,
	}
}
