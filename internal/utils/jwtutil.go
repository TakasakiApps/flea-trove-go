package utils

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hanakogo/exceptiongo"
	"time"
)

type JWTClaims[T any] struct {
	Data T
	jwt.RegisteredClaims
}

func JWTSign[T any](data T, key string, expireSecs int) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims[T]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expireSecs))),
		},
	})

	signedString, err := claims.SignedString([]byte(key))
	exceptiongo.ThrowErr[types.JWTSignError](err)

	return signedString
}

func mJWTParse[T any](tokenString string, key string) (*JWTClaims[T], bool) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims[T]{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if token == nil || token.Claims == nil {
		exceptiongo.ThrowErr[types.JWTParseError](err)
	}

	if jwtClaims, ok := token.Claims.(*JWTClaims[T]); ok {
		return jwtClaims, token.Valid
	}

	exceptiongo.ThrowErr[types.JWTParseError](err)
	return nil, false
}

func JWTParseData[T any](tokenString string, key string) T {
	claims, _ := mJWTParse[T](tokenString, key)
	return claims.Data
}

func JWTParse[T any](tokenString string, key string, callback func(isValid bool, data T)) {
	claims, isValid := mJWTParse[T](tokenString, key)
	if callback != nil {
		callback(isValid, claims.Data)
	}
}
