package repository

import (
	"github.com/supakjack/isekai-shop-api-tutorial/entities"
	_itemManagingModel "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/model"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error)
	Archiving(itemID uint64) error
}
