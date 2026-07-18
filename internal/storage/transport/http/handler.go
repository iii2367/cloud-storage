package http

import (
	"cloud-storage/internal/storage/service"
)

type Handler struct {
	service	*service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler {
		service: service,	
	}
}
