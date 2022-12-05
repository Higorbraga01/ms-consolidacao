package usecase

import (
	"context"
	"time"

	"github.com/higorbraga01/ms-cosolidacao/pkg/uow"
)

type MatchInput struct {
	ID      string
	Date    time.Time
	TeamAID string
	TeamBID string
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func (a *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {})
}
