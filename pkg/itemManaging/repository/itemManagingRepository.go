package repository

import (
	"github.com/MarkTBSS/084_Item_Archiving/entities"
	"github.com/MarkTBSS/084_Item_Archiving/pkg/itemManaging/models"
)

type ItemManagingRepository interface {
	Creating(itemEntity *entities.Item) (*entities.Item, error)
	Editing(itemID uint64, itemEditingReq *models.ItemEditingReq) (uint64, error)
	Archiving(itemID uint64) error
}
