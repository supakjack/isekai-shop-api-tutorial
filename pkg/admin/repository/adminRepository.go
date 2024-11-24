package repository

import "github.com/supakjack/isekai-shop-api-tutorial/entities"

type AdminRepository interface {
	Creating(playerEntity *entities.Admin) (*entities.Admin, error)
	FindById(playerId string) (*entities.Admin, error)
}
