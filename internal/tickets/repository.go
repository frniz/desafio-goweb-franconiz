package tickets

import (
	"fmt"

	"desafio-goweb-franconiz/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetByCountry(destination string) ([]domain.Ticket, error)
	GetAverageDestination(destination string) (float64, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetByCountry(destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetAverageDestination(destintation string) (float64, error) {

	tickets := r.db
	if len(tickets) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	var destinationCount int
	for _, t := range tickets {
		if t.Country == destintation {
			destinationCount++
		}
	}

	fmt.Printf("destinationCount: %v\n", destinationCount)
	fmt.Printf("len(tickets): %v\n", len(tickets))
	var average float64
	average = (float64(destinationCount) / float64(len(tickets))) * 100
	return average, nil
}
