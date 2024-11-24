package repository

import "github.com/supakjack/isekai-shop-api-tutorial/entities"

type PlayerRepository interface {
	Creating(playerEntity *entities.Player) (*entities.Player, error)
	FindById(playerID string) (*entities.Player, error)
}
