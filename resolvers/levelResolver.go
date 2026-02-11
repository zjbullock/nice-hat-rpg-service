package resolvers

import (
	"adventureBotService/models"
	"context"
)

type levelResolver struct {
	level models.Level
}

func (l *levelResolver) Value(_ context.Context) float64 {
	return float64(l.level.Value)
}

func (l *levelResolver) Exp(_ context.Context) float64 {
	return float64(l.level.Exp)
}
