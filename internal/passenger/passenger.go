package passenger

import (
	"bmc/bmc-assignment/internal/database"
)

type IService interface {
	Get(key string) (database.Passenger, error)
	All() ([]database.Passenger, error)
}

type passengerService struct {
	db database.IDatabase
}

func (p *passengerService) Get(key string) (database.Passenger, error) {
	return p.db.Get(key)
}

func (p *passengerService) All() ([]database.Passenger, error) {
	return p.db.All()
}

func New(db database.IDatabase) IService {
	return &passengerService{
		db: db,
	}
}
