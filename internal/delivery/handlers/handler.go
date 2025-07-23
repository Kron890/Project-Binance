package handlers

import "projectBinacne/internal/usecase"

type Handler struct {
	uc usecase.Ucecase
}

func NewHandler(uc usecase.Ucecase) *Handler {
	return &Handler{uc: uc}
}
