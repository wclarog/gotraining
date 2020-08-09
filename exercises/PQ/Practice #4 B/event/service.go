package event

import (
	"context"
	"github.com/go-kit/kit/log"
)

type Service interface {
	NextEvents(ctx context.Context) ([]Event, error)
}

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(r Repository, logger log.Logger) Service {
	return &service{
		repository: r,
		logger:     logger,
	}
}

func (s service) NextEvents(ctx context.Context) ([]Event, error) {
	result, err := s.repository.NextEvents(ctx)
	if err != nil {
		return []Event{}, err
	}

	return s.dtoToEventArray(result), nil
}

func (s service) dtoToEvent(f DTOEvent) Event {
	return Event{
		Title:       f.Title,
		Description: f.Description,
		Date:        f.Date,
	}
}

func (s service) dtoToEventArray(e []DTOEvent) []Event {
	var events []Event

	for _, current := range e {
		events = append(events, s.dtoToEvent(current) )
	}

	return events
}