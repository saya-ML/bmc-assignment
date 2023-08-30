package config

import (
	"bmc/bmc-assignment/internal/database"
	"bmc/bmc-assignment/internal/database/csv"
	"bmc/bmc-assignment/internal/database/sqlite"
	"bmc/bmc-assignment/internal/histogram"
	"bmc/bmc-assignment/internal/passenger"
	"os"
)

type Dependencies struct {
	PassengerService passenger.IService
	HistogramService histogram.IService
}

func BuildDependencies() Dependencies {
	source := os.Getenv("SOURCE")
	var db database.IDatabase
	if source == "sqlite3" {
		db = sqlite.New(os.Getenv("SQL_FILE"))
	} else {
		db = csv.New(os.Getenv("CSV_FILE"))
	}

	ps := passenger.New(db)
	return Dependencies{
		PassengerService: ps,
		HistogramService: histogram.New(ps),
	}
}
