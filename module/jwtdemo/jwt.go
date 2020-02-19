package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"time"
)

/**
	1.自定义payload结构体,不建议直接使用 dgrijalva/jwt-go jwt.StandardClaims结构体.因为他的payload包含的用户信息太少.
	2.实现 type Claims interface 的 Valid() error 方法,自定义校验内容
	3_工厂方法.生成JWT-string jwtGenerateToken(m *User,d time.Duration) (string, error)
 */

type User struct {
	ID        int    `json:"id"`
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
	DDingId  string `json:"dd_id"`

}

var AppSecret = ""                      //viper.GetString会设置这个值(32byte长度)
var AppIss = "forchange/rebind"         //这个值会被viper.GetString重写


type userClaims struct {
	jwt.StandardClaims
	*User
}

//实现 `type Claims interface` 的 `Valid() error` 方法,自定义校验内容
func (c userClaims) Valid() (err error) {
	// 校验是否过期
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return  errors.New("token is expired")
	}
	// 校验AppIss
	if !c.VerifyIssuer(AppIss, true) {
		return  errors.New("token's issuer is wrong")
	}

	return
}


// 生成token
func JwtGenerateToken(u *User, d time.Duration) (string, error) {

	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%s", u.ID),
		Issuer:    AppIss,
	}
	uClaims := userClaims{
		StandardClaims: stdClaims,
		User:           u,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		log.Fatal().Msg("config is wrong, can not generate jwt")
	}
	return tokenString, err
}


//JwtParseUser 解析payload的内容,得到用户信息
func JwtParseUser(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := userClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.User, err
}

func main() {
	//u := &User{
	//	ID:     30,
	//	Mobile: "18860107639",
	//	DDingId:"12312312423131",
	//	UserName:"fengzhiwen",
	//}
	//token, err := JwtGenerateToken(u, time.Minute*5)
	//if err != nil {
	//	log.Fatal().Msg(err.Error())
	//}
	//fmt.Println(token)

	mobile, err := JwtParseUser("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJleHAiOjE1ODEzMzk3MTgsImp0aSI6IiUhcyhpbnQ9MzApIiwiaWF0IjoxNTgxMzM5NDE4LCJpc3MiOiJmb3JjaGFuZ2UvcmViaW5kIiwi" +
		"aWQiOjMwLCJtb2JpbGUiOiIxODg2MDEwNzYzOSIsInVzZXJfbmFtZSI6ImZlbmd6aGl3ZW4iLCJkZF9pZCI6IjEyMzEyMzEyNDIzMTMxIn0._wdqJnXGGaQptRjptPEAl-rvCnNhC2l-dkKGoZGGxsA")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	fmt.Println(mobile)
}
