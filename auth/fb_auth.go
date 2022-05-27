package auth

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
)

// FBV4 firebase v4 admin implementation of IAuth
type FBV4 struct {
	client *auth.Client
}

// VerifyToken validates refresh token and returns uid of the requestor
func (fa *FBV4) VerifyToken(ctx context.Context, token string) (string, error) {

	ut, err := fa.client.VerifyIDToken(ctx, token)
	if err != nil {
		return "", fmt.Errorf("failed to verify ID token: %v", err)
	}

	return ut.UID, nil
}

// RevokeToken revokes all refresh tokens requested by the requestor
func (fa *FBV4) RevokeToken(ctx context.Context, uid string) error {
	return fa.client.RevokeRefreshTokens(ctx, uid)
}
