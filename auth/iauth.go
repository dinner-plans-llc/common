package auth

import "context"

// IAuth
type IAuth interface {
	VerifyToken(ctx context.Context, token string) (string, error)
	RevokeToken(ctx context.Context, uid string) error
}
