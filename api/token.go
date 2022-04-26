package api

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/Model"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID  string `json:"name"`
	Pwd     string `json:"pwd"`
	IsAdmin bool   `json:"isAdmin"`
}

var (
	Secret     = "fsruav" //密钥
	ExpireTime = 3600     //token的有效期
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登录"
)

// getToken token获取
func getToken(claims *JWTClaims) (string, error) {
	//token获取
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//token密钥加密
	signedToken, err := token.SignedString([]byte(Secret))

	//加密失败
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}

	//成功返回token
	return signedToken, nil
}

// Login 登录
func Login(c *gin.Context) {
	/*
		检查用户是否存在
		判断密码是否正确
		发放token给前端
	*/

	//登录结构体定义
	type UserLogin struct {
		UserID  string `json:"stuid"`
		Pwd     string `json:"pwd"`
		IsAdmin bool   `json:"isAdmin"`
	}
	var user UserLogin

	//绑定结构体
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"code": 400, "desc": "上传格式出错"})
		return
	}

	//用户验证
	User, err := Model.GetUserByIDToLogin(user.UserID)
	//用户不存在
	if err != nil {
		c.JSON(404, gin.H{"code": 404, "desc": "未找到该用户"})
		return
	}
	//密码验证
	if User.Pwd != user.Pwd {
		c.JSON(401, gin.H{"code": 401, "desc": "密码错误"})
	}

	//claims初始化
	claims := &JWTClaims{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	}
	//发布时间
	claims.IssuedAt = time.Now().Unix()
	//有效时间
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	//token获取
	signedToken, err := getToken(claims)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{"token": signedToken, "desc": "登录成功", "IsAdmin": User.IsAdmin})
}

// verifyAction
func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	clamis, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify")
	return clamis, nil
}

// AuthRequired 验证token
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken := c.Param("token")
		claim, err := verifyAction(strToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
		}
		c.JSON(200, gin.H{"desc": claim.UserID + "verify"})
		c.Next()
	}
}
