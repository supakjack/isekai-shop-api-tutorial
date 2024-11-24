package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/supakjack/isekai-shop-api-tutorial/databases"
	"github.com/supakjack/isekai-shop-api-tutorial/entities"
	_itemManagingException "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/exception"
	_itemManagingModel "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/model"
)

type itemManagingRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewItemManagingRepositoryImpl(db databases.Database, logger echo.Logger) ItemManagingRepository {
	return &itemManagingRepositoryImpl{db, logger}
}

func (r *itemManagingRepositoryImpl) Creating(itemEntity *entities.Item) (*entities.Item, error) {
	item := new(entities.Item)
	err := r.db.Connect().Create(itemEntity).Scan(item).Error

	if err != nil {
		r.logger.Errorf("Creating item failed: %s", err.Error())
		return nil, err
	}

	return item, nil
}

func (r *itemManagingRepositoryImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (uint64, error) {
	err := r.db.Connect().Model(&entities.Item{}).Where("id = ?", itemID).Updates(itemEditingReq).Error

	if err != nil {
		r.logger.Error("Editing item failed %s", err.Error())
		return 0, &_itemManagingException.ItemEditing{}
	}

	return itemID, nil
}

func (r *itemManagingRepositoryImpl) Archiving(itemID uint64) error {
	err := r.db.Connect().Table("items").Where("id = ?", itemID).Update("is_archive", true).Error
	if err != nil {
		r.logger.Errorf("Archiving item failed: %s", err.Error())
		return &_itemManagingException.ItemArchiving{ItemID: itemID}
	}

	return nil
}
