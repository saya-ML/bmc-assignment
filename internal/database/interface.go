package database

import "reflect"

type IDatabase interface {
	Get(key string) (Passenger, error)
	All() ([]Passenger, error)
}

type Passenger struct {
	PassengerId string  `json:"passenger_id"`
	Survived    string  `json:"survived"`
	Pclass      string  `json:"pclass"`
	Name        string  `json:"name"`
	Sex         string  `json:"sex"`
	Age         string  `json:"age"`
	SibSp       string  `json:"sibSp"`
	Parch       string  `json:"parch"`
	Ticket      string  `json:"ticket"`
	Fare        float64 `json:"fare"`
	Cabin       string  `json:"cabin"`
	Embarked    string  `json:"embarked"`
}

func (s *Passenger) SelectFields(fields []string) map[string]interface{} {
	fs := fieldSet(fields...)
	rt, rv := reflect.TypeOf(*s), reflect.ValueOf(*s)
	out := make(map[string]interface{}, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")
		if _, found := fs[jsonKey]; found {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}
	return out
}

func fieldSet(fields ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(fields))
	for _, s := range fields {
		set[s] = struct{}{}
	}
	return set
}
