package task4

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type myClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   uint   `json:"user_id"`
	jwt.StandardClaims
}

const TokenExpireTime = time.Hour * 2

var JwtSecret interface{} = []byte("Intex01SecretKey")

func GenerateJWT(user User) (string, error) {
	// Implement JWT generation logic here
	claims := myClaims{
		Username: user.Username,
		Email:    user.Email,
		UserId:   user.ID,
	}
	claims.ExpiresAt = time.Now().Add(TokenExpireTime).Unix()
	claims.Issuer = "Intex01"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ParseToken(tokenString string) (*myClaims, error) {
	//解析token
	//jwt.ParseWithClaims 方法用于解析鉴权的声明，返回 *jwt.Token。
	token, err := jwt.ParseWithClaims(tokenString, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	//Valid 方法用于校验鉴权的声明
	if claims, ok := token.Claims.(*myClaims); ok && token.Valid { //校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("email", mc.Email)
		c.Set("user_id", mc.UserId)
		c.Next()
	}
}

func GetUserIDFromContext(c *gin.Context) (uint, error) {
	v, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("user_id not found in context")
	}
	switch id := v.(type) {
	case uint:
		return id, nil
	case int:
		return uint(id), nil
	case int64:
		return uint(id), nil
	case float64:
		return uint(id), nil
	case string:
		ui, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(ui), nil
	default:
		return 0, fmt.Errorf("unsupported user_id type %T", v)
	}
}
