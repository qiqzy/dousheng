package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/common"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//聚合jwt- 内部实现的Claims
type CustomerClaim struct {
	common.User
	*jwt.StandardClaims
}

func Encode(privateKey string, user common.User) (string, error) {
	//设置超时时间
	expTime := time.Now().Add(time.Hour * 24 * 3).Unix()

	//设置Claim
	customer := CustomerClaim{user, &jwt.StandardClaims{ExpiresAt: expTime}}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customer)

	return token.SignedString([]byte(privateKey))
}

//token字符串解码成用户信息
func Decode(privateKey string, tokenString string) (*common.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaim{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(privateKey), nil
	})

	if err != nil {
		fmt.Println("token解码出错: ", err, "接收到的token为：", tokenString)
		return nil, err
	}

	if claim, ok := token.Claims.(*CustomerClaim); ok && token.Valid {
		return &claim.User, nil
	} else {
		return nil, err
	}
}
