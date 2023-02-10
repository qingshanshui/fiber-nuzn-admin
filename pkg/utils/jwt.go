package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GetToken(name, uid string) (string, error) {
	// 加密内容信息
	/**
	iss (issuer)：签发人
	exp (expiration time)：过期时间
	sub (subject)：主题
	aud (audience)：受众
	nbf (Not Before)：生效时间
	iat (Issued At)：签发时间
	jti (JWT ID)：编号
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"uid":  uid,
		"exp":  time.Now().Add(time.Second * time.Duration(viper.GetInt("Jwt.Expire"))).Unix(),
		"iss":  "nuzn.cn",
	})
	tokenString, err := token.SignedString([]byte(viper.GetString("Jwt.Secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	//在这里如果也使用jwt.ParseWithClaims的话，第二个参数就写jwt.MapClaims{}
	//例如jwt.ParseWithClaims(tokenString, jwt.MapClaims{},func(t *jwt.Token) (interface{}, error){}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("Jwt.Secret")), nil
	})
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	fmt.Println(claims["foo"], claims["nbf"])
	//} else {
	//	fmt.Println(err)
	//}
	// 检查token是否合法
	if token.Valid {
		return token, nil
	}
	return nil, err
}
