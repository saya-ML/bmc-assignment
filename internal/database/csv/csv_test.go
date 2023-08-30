package csv

import (
	"bmc/bmc-assignment/internal/database"
	"errors"
	"reflect"
	"testing"
)

const CSVFile3Rows = "csv.input"
const CSVInvalidFare = "csv.invalidFare.input"
const CSVEmpty = "csv.empty.input"

func Test_dbCSV_GetPassenger(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		key      string
		want     database.Passenger
		wantErr  bool
		err      error
	}{
		{
			name:     "file not found",
			filename: "notFoundFile",
			wantErr:  true,
			err:      errors.New("open notFoundFile: no such file or directory"),
		},
		{
			name:     "key not found",
			filename: CSVFile3Rows,
			key:      "keyNotFound",
			want:     database.Passenger{},
			wantErr:  true,
			err:      errors.New("passenger not found"),
		},
		{
			name:     "key found",
			filename: CSVFile3Rows,
			key:      "25",
			want:     passenger25(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(tt.filename)
			got, err := d.Get(tt.key)
			if tt.wantErr && err.Error() != tt.err.Error() {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recordToDBDto(t *testing.T) {
	tests := []struct {
		name    string
		record  []string
		want    database.Passenger
		wantErr bool
	}{
		{
			name:    "invalid fare",
			record:  []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "NotANumber"},
			want:    database.Passenger{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := recordToDBDto(tt.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("recordToDBDto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recordToDBDto() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dbCSV_GetPassengerAll(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []database.Passenger
		wantErr  bool
		err      error
	}{
		{
			name:     "Get all invalid data",
			filename: CSVInvalidFare,
			want:     nil,
			wantErr:  true,
			err:      errors.New("strconv.ParseFloat: parsing \"8A.4583\": invalid syntax"),
		},
		{
			name:     "Get all",
			filename: CSVFile3Rows,
			want:     dbData(),
		},
		{
			name:     "Get empty file",
			filename: CSVEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbCSV{
				filename: tt.filename,
			}
			got, err := d.All()
			if tt.wantErr && err.Error() != tt.err.Error() {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func passenger25() database.Passenger {
	return database.Passenger{
		PassengerId: "25",
		Survived:    "0",
		Pclass:      "3",
		Name:        "Palsson, Miss. Torborg Danira",
		Sex:         "female",
		Age:         "8",
		SibSp:       "3",
		Parch:       "1",
		Ticket:      "349909",
		Fare:        21.075,
		Cabin:       "",
		Embarked:    "S",
	}
}

func dbData() []database.Passenger {
	return []database.Passenger{
		{
			PassengerId: "22",
			Survived:    "1",
			Pclass:      "2",
			Name:        "Beesley, Mr. Lawrence",
			Sex:         "male",
			Age:         "34",
			SibSp:       "0",
			Parch:       "0",
			Ticket:      "248698",
			Fare:        13,
			Cabin:       "D56",
			Embarked:    "S",
		},
		passenger25(),
		{
			PassengerId: "854",
			Survived:    "1",
			Pclass:      "1",
			Name:        "Lines, Miss. Mary Conover",
			Sex:         "female",
			Age:         "16",
			SibSp:       "0",
			Parch:       "1",
			Ticket:      "PC 17592",
			Fare:        39.4,
			Cabin:       "D28",
			Embarked:    "S",
		},
	}
}
