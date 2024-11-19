package repository

import "github.com/supakjack/isekai-shop-api-tutorial/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
