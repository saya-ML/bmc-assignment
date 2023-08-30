package handlers

import (
	"bmc/bmc-assignment/internal/database"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func Test_passengerHandler_GetPassenger(t *testing.T) {
	tests := []struct {
		name      string
		passenger database.Passenger
		params    url.Values
		err       error
	}{
		{
			name:      "get ok",
			passenger: passenger4855(),
			err:       nil,
		},
		{
			name:      "get error",
			passenger: database.Passenger{},
			err:       errors.New("generic error"),
		},
		{
			name:      "get filtered",
			passenger: passenger4855(),
			params:    url.Values{"fields": []string{"passenger_id", "fare", "cabin"}},
			err:       nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var errNewRequest error
			mockPassengerService := NewPassengerMock()
			mockPassengerService.On("Get").Return(tt.passenger, tt.err).Once()
			p := NewPassenger(mockPassengerService)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request, errNewRequest = http.NewRequest(http.MethodGet, "/?"+tt.params.Encode(), nil)
			assert.Nil(t, errNewRequest)
			p.Get(ctx)

			if tt.err == nil {
				assert.Equal(t, w.Code, http.StatusOK)
				var passenger database.Passenger
				unmarshalError := json.Unmarshal(w.Body.Bytes(), &passenger)
				assert.Nil(t, unmarshalError)
				if len(tt.params) == 0 {
					if !reflect.DeepEqual(passenger, tt.passenger) {
						t.Errorf("Get() got = %v, want %v", passenger, tt.passenger)
					}
				}
			} else {
				assert.Equal(t, w.Code, http.StatusBadRequest)
				errMessage := map[string]string{}
				errUnmarshal := json.Unmarshal(w.Body.Bytes(), &errMessage)
				assert.Nil(t, errUnmarshal)
				got := errMessage["error"]
				if got != tt.err.Error() {
					t.Errorf("Get() error = %v, want %v", got, tt.err.Error())
				}
			}

		})
	}
}

func Test_passengerHandler_GetPassengerAll(t *testing.T) {
	tests := []struct {
		name       string
		passengers []database.Passenger
		err        error
	}{
		{
			name:       "get all OK",
			passengers: dbData(),
			err:        nil,
		},
		{
			name:       "get all error",
			passengers: nil,
			err:        errors.New("new error to view"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPassengerService := NewPassengerMock()
			mockPassengerService.On("All").Return(tt.passengers, tt.err).Once()
			p := NewPassenger(mockPassengerService)
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			p.All(ctx)

			if tt.err == nil {
				assert.Equal(t, w.Code, http.StatusOK)
				var passengers []database.Passenger
				unmarshalError := json.Unmarshal(w.Body.Bytes(), &passengers)
				assert.Nil(t, unmarshalError)
				if !reflect.DeepEqual(passengers, tt.passengers) {
					t.Errorf("Get() got = %v, want %v", passengers, tt.passengers)
				}
			} else {
				assert.Equal(t, w.Code, http.StatusBadRequest)
				errMessage := map[string]string{}
				errUnmarshal := json.Unmarshal(w.Body.Bytes(), &errMessage)
				assert.Nil(t, errUnmarshal)
				got := errMessage["error"]
				if got != tt.err.Error() {
					t.Errorf("Get() error = %v, want %v", got, tt.err.Error())
				}
			}
		})
	}
}

type MockPassengerService struct {
	mock.Mock
}

func NewPassengerMock() *MockPassengerService {
	return new(MockPassengerService)
}

func (m *MockPassengerService) Get(_ string) (database.Passenger, error) {
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

func (m *MockPassengerService) All() ([]database.Passenger, error) {
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

func passenger4855() database.Passenger {
	return database.Passenger{
		PassengerId: "4855",
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
			PassengerId: "12",
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
		passenger4855(),
		{
			PassengerId: "4856",
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
