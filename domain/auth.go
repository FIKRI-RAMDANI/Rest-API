package domain

import (
	"context"

	"github.com/FIKRI-RAMDANI/Rest-API/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
