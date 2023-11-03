package utils

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/hanakogo/exceptiongo"
	"golang.org/x/crypto/bcrypt"
)

func BCryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	exceptiongo.ThrowErr[types.CryptError](err)
	return string(hash)
}

func CompareP2B(password string, bcryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptPwd), []byte(password))
	return err == nil
}
