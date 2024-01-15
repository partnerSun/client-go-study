package jwtutils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goStudy/config"
	"goStudy/utils/logs"
	"time"
)

var mySigningKey = []byte(config.SigningKey)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// token生成函数
func Gentoken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.JwtExpireTime))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Personal",
			Subject:   "partnerSun",
			ID:        "1",
		},
	}
	//println("mySigningKey:", mySigningKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	//println("token:", token)
	//println("ss:", ss)
	return ss, err
}

// token解析函数
func ParseJwtToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
	if err != nil {
		logs.Error(nil, "token解析失败")
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("\n用户是: %v \n发行机构是: %v \n过期时间是: %v", claims.Username, claims.RegisteredClaims.Issuer, claims.RegisteredClaims.ExpiresAt)
		return claims, nil

	} else {
		logs.Warning(nil, "token解析失败")
		return nil, errors.New("token不合法，失效的token")

	}
}
