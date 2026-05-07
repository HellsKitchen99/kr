package http

import "github.com/HellsKitchen99/kr/app/internal/domain"

type Service interface {
	GetInfo() domain.Info
}
