package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakjack/isekai-shop-api-tutorial/databases"
	"github.com/supakjack/isekai-shop-api-tutorial/entities"
	_playerException "github.com/supakjack/isekai-shop-api-tutorial/pkg/player/exception"
)

type playerRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewPlayerRepositoryImpl(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{db, logger}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	player := new(entities.Player)
	err := r.db.Connect().Create(playerEntity).Scan(player).Error

	if err != nil {
		r.logger.Errorf("Creating player failed: %s", err.Error())
		return nil, &_playerException.PlayerCreating{PlayerID: playerEntity.ID}
	}

	return player, nil
}

func (r *playerRepositoryImpl) FindById(playerID string) (*entities.Player, error) {
	player := new(entities.Player)

	err := r.db.Connect().First(player, playerID).Error
	if err != nil {
		r.logger.Errorf("Failed to find admin by adminId : %s", err.Error())
		return nil, &_playerException.PlayerNotFound{PlayerID: player.ID}
	}

	return player, nil
}
