package passenger

import (
	"bmc/bmc-assignment/internal/database"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_passengerService_GetPassenger(t *testing.T) {

	tests := []struct {
		name    string
		key     string
		want    database.Passenger
		wantErr bool
	}{
		{
			name:    "key 854",
			key:     "854",
			want:    passenger855(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewDBMock()
			db.On("Get").Return(passenger855(), nil).Once()
			p := New(db)
			got, err := p.Get(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passengerService_GetPassengerAll(t *testing.T) {
	tests := []struct {
		name    string
		want    []database.Passenger
		wantErr bool
	}{
		{
			name:    "Get all SVC",
			want:    dbData(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := NewDBMock()
			db.On("All").Return(dbData(), nil).Once()
			p := New(db)
			got, err := p.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func passenger855() database.Passenger {
	return database.Passenger{
		PassengerId: "855",
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
			PassengerId: "21",
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
		passenger855(),
		{
			PassengerId: "856",
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

type MockDBService struct {
	mock.Mock
}

func NewDBMock() *MockDBService {
	return new(MockDBService)
}

func (m *MockDBService) Get(key string) (database.Passenger, error) {
	ret := m.Called()
	var passenger database.Passenger
	if ret.Get(0) != nil {
		passenger = ret.Get(0).(database.Passenger)
	}
	var err error
	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}
	return passenger, err
}

func (m *MockDBService) All() ([]database.Passenger, error) {
	ret := m.Called()
	var passengers []database.Passenger
	if ret.Get(0) != nil {
		passengers = ret.Get(0).([]database.Passenger)
	}
	var err error
	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}
	return passengers, err
}
