package mid

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func StreamAuthInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		md, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		token := md["authorization"]
		if len(token) == 0 {
			return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		var tokenString string
		if len(token) > 0 {
			tokenParts := strings.SplitN(token[0], " ", 2)
			if len(tokenParts) == 2 && strings.ToLower(tokenParts[0]) == "bearer" {
				tokenString = tokenParts[1]
			} else {
				return status.Errorf(codes.Unauthenticated, "invalid authorization header format")
			}
		} else {
			return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		secretKey := []byte("grilsprower")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			return status.Errorf(codes.Unauthenticated, "invalid authorization token: %v", err)
		}

		return handler(srv, stream)
	}
}
