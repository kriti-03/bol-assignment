package service

import (
	"github.com/pablocrivella/mancala/internal/datastore"
	"github.com/pablocrivella/mancala/internal/engine"
)

// MancalaService is a games context domain service
type MancalaService struct {
	gameRepo datastore.GameRepo
}

// NewService returns a games.MancalaService
func NewService(g datastore.GameRepo) MancalaService {
	return MancalaService{gameRepo: g}
}

func (s MancalaService) CreateGame(player1, player2 string) (engine.Game, error) {
	g := engine.NewBoard(player1, player2)
	if err := s.gameRepo.Save(g); err != nil {
		return engine.Game{}, err
	}
	return g, nil
}

func (s MancalaService) FindGame(id string) (engine.Game, error) {
	g, err := s.gameRepo.Find(id)
	if err != nil {
		return engine.Game{}, err
	}
	return g, nil
}

func (s MancalaService) ExecutePlay(gameID string, pitIndex int64) (engine.Game, error) {
	g, err := s.gameRepo.Find(gameID)
	if err != nil {
		return engine.Game{}, err
	}

	err = g.PlayTurn(pitIndex)
	if err != nil {
		return engine.Game{}, err
	}

	err = s.gameRepo.Save(g)
	if err != nil {
		return engine.Game{}, err
	}
	return g, nil
}
