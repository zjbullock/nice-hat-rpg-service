package resolvers

import (
	"adventureBotService/models"
	"context"
)

type areaResponseResolver struct {
	areaInfo *models.Area
	message  *string
}

func (a *areaResponseResolver) AreaInfo(_ context.Context) *areaResolver {
	return &areaResolver{area: *a.areaInfo}
}

func (a *areaResponseResolver) Message(_ context.Context) *string {
	return a.message
}
