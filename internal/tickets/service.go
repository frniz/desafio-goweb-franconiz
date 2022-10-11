package tickets

import (
	"desafio-goweb-franconiz/internal/domain"
)

type Service interface {
	GetTotalTickets() ([]domain.Ticket, error)
	GetByCountry(destination string) ([]domain.Ticket, error)
	GetAverageDestination(destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets() ([]domain.Ticket, error) {

	rep, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return rep, nil
}

func (s *service) GetByCountry(destination string) ([]domain.Ticket, error) {

	tickets, err := s.repository.GetByCountry(destination)
	if err != nil {

		return nil, err
	}
	return tickets, nil
}

func (s *service) GetAverageDestination(destintation string) (float64, error) {

	average, err := s.repository.GetAverageDestination(destintation)
	if err != nil {
		return 0, err
	}
	return average, nil
}
