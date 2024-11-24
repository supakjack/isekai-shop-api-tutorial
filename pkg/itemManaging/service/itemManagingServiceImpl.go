package service

import (
	"github.com/supakjack/isekai-shop-api-tutorial/entities"
	_itemManagingModel "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/model"
	_itemManagingRepository "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/repository"
	_itemShopModel "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemShop/model"
	_itemShopRepository "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemShop/repository"
)

type itemManagingServiceImpl struct {
	itemManagingRepository _itemManagingRepository.ItemManagingRepository
	itemShopRepository     _itemShopRepository.ItemShopRepository
}

func NewItemManagingServiceImpl(
	itemManagingRepository _itemManagingRepository.ItemManagingRepository,
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemManagingService {
	return &itemManagingServiceImpl{itemManagingRepository, itemShopRepository}
}

func (s *itemManagingServiceImpl) Creating(itemCreatingReq *_itemManagingModel.ItemCreatingReq) (*_itemShopModel.Item, error) {
	itemEntity := &entities.Item{
		Name:        itemCreatingReq.Name,
		Description: itemCreatingReq.Description,
		Picture:     itemCreatingReq.Picture,
		Price:       itemCreatingReq.Price,
	}

	itemEntityResult, err := s.itemManagingRepository.Creating(itemEntity)

	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Editing(itemID uint64, itemEditingReq *_itemManagingModel.ItemEditingReq) (*_itemShopModel.Item, error) {
	_, err := s.itemManagingRepository.Editing(itemID, itemEditingReq)
	if err != nil {
		return nil, err
	}

	itemEntityResult, err := s.itemShopRepository.FindById(itemID)
	if err != nil {
		return nil, err
	}

	return itemEntityResult.ToItemModel(), nil
}

func (s *itemManagingServiceImpl) Archiving(itemID uint64) error {
	return s.itemManagingRepository.Archiving(itemID)
}
