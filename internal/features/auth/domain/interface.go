package domain

import (
	"context"

	roleDomain "github.com/ramdhanrizkij/arastore-api/internal/features/role/domain"
	userDomain "github.com/ramdhanrizkij/arastore-api/internal/features/user/domain"
)

// AuthRepository defines the data-access contract for the auth feature.
// Implementations live in the repository layer and use GORM.
type AuthRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*userDomain.User, error)
	CreateUser(ctx context.Context, user *userDomain.User) error
	FindRoleByName(ctx context.Context, name string) (*roleDomain.Role, error)
	CreateRefreshToken(ctx context.Context, token *RefreshToken) error
	FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, id string) error
	CleanupExpiredTokens(ctx context.Context) error
}

// AuthService defines the business-logic contract for the auth feature.
// Implementations are transport-agnostic (no Fiber / gRPC).
type AuthService interface {
	Register(ctx context.Context, req *RegisterRequest) (*AuthResponse, error)
	Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error)
	Refresh(ctx context.Context, req *RefreshTokenRequest) (*TokenResponse, error)
	Logout(ctx context.Context, req *LogoutRequest) error
	CleanupExpiredTokens(ctx context.Context) error
}
