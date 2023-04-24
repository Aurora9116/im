package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var myKey = []byte("im")

type UserClaims struct {
	Identity string
	Email    string
	jwt.StandardClaims
}

func Md5(s string) string {
	hash := md5.New()
	return fmt.Sprintf("%x", hash.Sum([]byte(s)))
}

func GenerateToken(identity, email string) (string, error) {
	claim := &UserClaims{
		Identity:       identity,
		Email:          email,
		StandardClaims: jwt.StandardClaims{}, // 暂时未做刷新处理
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	claims, err := jwt.ParseWithClaims(tokenString, new(UserClaims), func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := claims.Claims.(*UserClaims); ok && claims.Valid {
		return claim, nil
	}
	return nil, errors.New("token错误")
}
