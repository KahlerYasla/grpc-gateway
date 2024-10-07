package auth

import (
	"context"
	"time"

	"/src/internal/service/auth/proto/gen"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const secretKey = "grilsprower"

type AuthService struct {
	gen.UnimplementedAuthServiceServer
}

func (s *AuthService) GenerateToken(ctx context.Context, req *gen.AuthRequest) (*gen.AuthResponse, error) {
	// Perform your authentication logic here
	if req.Username == "user" && req.Password == "pass" { // Todo: Replace with actual authentication
		token, err := generateJWT(req.Username)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not generate token: %v", err)
		}
		return &gen.AuthResponse{Token: token}, nil
	}
	return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
}

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
