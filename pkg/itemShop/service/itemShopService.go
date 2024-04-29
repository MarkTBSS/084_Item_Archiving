package service

import _models "github.com/MarkTBSS/084_Item_Archiving/pkg/itemShop/models"

type ItemShopService interface {
	Listing(itemFilter *_models.ItemFilter) (*_models.ItemResult, error)
}
