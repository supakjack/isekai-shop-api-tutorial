package service

import (
	_itemShopModel "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemShop/model"
	_itemShopRepository "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(
	itemShopRepository _itemShopRepository.ItemShopRepository,
) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}

func (s *itemShopServiceImpl) Listing() ([]*_itemShopModel.Item, error) {
	itemList, err := s.itemShopRepository.Listing()
	if err != nil {
		return nil, err
	}

	itemModelList := make([]*_itemShopModel.Item, 0)

	for _, item := range itemList {
		itemModelList = append(itemModelList, item.ToItemModel())
	}

	return itemModelList, nil
}
