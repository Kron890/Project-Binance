package usecase

import (
	binanceapi "projectBinacne/internal/repository/binance_api"
	repository "projectBinacne/internal/repository/postgres"
)

type Ucecase struct {
	Repo       *repository.PostgresRepo
	BinService *binanceapi.BinanceService
}

func NewUsecase(r *repository.PostgresRepo, b *binanceapi.BinanceService) *Ucecase {
	return &Ucecase{
		Repo:       r,
		BinService: b}

}
