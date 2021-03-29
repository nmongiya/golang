package service

import (
	"entain-master/racing/db"
	"entain-master/racing/proto/racing"
	"golang.org/x/net/context"
)

type Racing interface {
	// ListRaces will return a collection of races.
	ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error)
	GetRace(ctx context.Context, in *racing.GetRaceRequest) (*racing.ListRacesResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

func (s *racingService) GetRace(ctx context.Context, in *racing.GetRaceRequest) (*racing.ListRacesResponse, error) {
    races, err := s.racesRepo.GetRace(in.GetId())
	if err != nil {
		return nil, err
	}

	return &racing.ListRacesResponse{Races: races}, nil

}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
	races, err := s.racesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &racing.ListRacesResponse{Races: races}, nil
}
