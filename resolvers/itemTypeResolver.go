package resolvers

import (
	"adventureBotService/models"
	"context"
)

type itemTypeResolver struct {
	itemType *models.ItemType
}

func (i *itemTypeResolver) Type(_ context.Context) string {
	return i.itemType.Type
}

func (i *itemTypeResolver) WeaponType(_ context.Context) *string {
	return i.itemType.WeaponType
}
