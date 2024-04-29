package repository

import (
	_entities "github.com/MarkTBSS/084_Item_Archiving/entities"
	_models "github.com/MarkTBSS/084_Item_Archiving/pkg/itemShop/models"
)

type ItemShopRepository interface {
	Listing(itemFilter *_models.ItemFilter) ([]*_entities.Item, error)
	Counting(itemFilterDto *_models.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*_entities.Item, error)
}
