package server

import (
	_itemManagingController "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/controller"
	_itemManagingRepository "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/repository"
	_itemManagingService "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemManaging/service"
	_itemShopRepository "github.com/supakjack/isekai-shop-api-tutorial/pkg/itemShop/repository"
)

func (s *echoServer) initItemManagingRouter() {
	router := s.app.Group("/v1/item-managing")

	itemRepository := _itemShopRepository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemMangingRepository := _itemManagingRepository.NewItemManagingRepositoryImpl(s.db, s.app.Logger)
	itemManagingService := _itemManagingService.NewItemManagingServiceImpl(itemMangingRepository, itemRepository)

	itemManagingController := _itemManagingController.NewItemManagingControllerImpl(itemManagingService)

	router.POST("", itemManagingController.Creating)
	router.PATCH("/:itemID", itemManagingController.Editing)
	router.DELETE("/:itemID", itemManagingController.Archiving)
}
