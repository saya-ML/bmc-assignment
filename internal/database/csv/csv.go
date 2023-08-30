package csv

import (
	"bmc/bmc-assignment/internal/database"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

type PassengerCSV struct {
	PassengerId string `csv:"PassengerId"`
	Survived    string `csv:"Survived"`
	Pclass      string `csv:"Pclass"`
	Name        string `csv:"Name"`
	Sex         string `csv:"Sex"`
	Age         string `csv:"Age"`
	SibSp       string `csv:"SibSp"`
	Parch       string `csv:"Parch"`
	Ticket      string `csv:"Ticket"`
	Fare        string `csv:"Fare"`
	Cabin       string `csv:"Cabin"`
	Embarked    string `csv:"Embarked"`
}

type dbCSV struct {
	filename string
}

func New(filename string) database.IDatabase {
	return &dbCSV{
		filename: filename,
	}
}

func (d *dbCSV) Get(key string) (database.Passenger, error) {
	p, err := d.scanFile(&key)
	if err != nil {
		return database.Passenger{}, err
	}
	if len(p) == 0 {
		return database.Passenger{}, errors.New("passenger not found")
	}
	return p[0], nil
}

func (d *dbCSV) All() ([]database.Passenger, error) {
	return d.scanFile(nil)
}

func (d *dbCSV) scanFile(key *string) ([]database.Passenger, error) {
	f, err := os.Open(d.filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var passengers []database.Passenger
	r := csv.NewReader(f)
	var record []string
	var currentPassenger database.Passenger
	_, err = r.Read() // skip headers
	for err == nil {
		record, err = r.Read()
		if err != nil {
			break
		}
		currentPassenger, err = recordToDBDto(record)
		if err != nil {
			return nil, err
		}
		if key == nil || *key == record[0] {
			passengers = append(passengers, currentPassenger)
			if key != nil {
				return passengers, nil
			}
		}
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return passengers, nil
}

func recordToDBDto(record []string) (database.Passenger, error) {
	fare, err := strconv.ParseFloat(record[9], 64)
	if err != nil {
		return database.Passenger{}, err
	}
	return database.Passenger{
		PassengerId: record[0],
		Survived:    record[1],
		Pclass:      record[2],
		Name:        record[3],
		Sex:         record[4],
		Age:         record[5],
		SibSp:       record[6],
		Parch:       record[7],
		Ticket:      record[8],
		Fare:        fare,
		Cabin:       record[10],
		Embarked:    record[11],
	}, nil
}
