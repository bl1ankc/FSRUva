package Mid

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/Controller"
	"main/Service"
	"main/Service/Status"
	"main/utils"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID    string `json:"name"`
	Pwd       string `json:"pwd"`
	IsAdmin   bool   `json:"isAdmin"`
	AdminType int    `json:"adminType"`
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
		UserID    string `json:"stuid"`
		Pwd       string `json:"pwd"`
		NickName  string `json:"nickName"`
		AvatarUrl string `json:"avatarUrl"`
		Code      string `json:"code"`
		//IsAdmin bool   `json:"isAdmin"`
	}
	var user UserLogin

	//绑定结构体
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"code": 400, "desc": "上传格式出错"})
		return
	}

	//用户验证
	User, err := Service.GetUserByIDToLogin(user.UserID)
	//用户不存在
	if err != nil {
		c.JSON(404, gin.H{"code": 404, "desc": "未找到该用户"})
		return
	}
	//加密
	h := md5.New()
	h.Write([]byte(user.Pwd))
	ciphertext := hex.EncodeToString(h.Sum(nil))

	//密码验证
	if User.Pwd != ciphertext {
		c.JSON(401, gin.H{"code": 401, "desc": "密码错误"})
		return
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
	_, Openid, Unionid := utils.GetOpenid(user.Code)
	Service.UpdateUserInfo(user.UserID, user.NickName, user.AvatarUrl, Openid, Unionid)
	c.JSON(200, gin.H{"token": signedToken, "desc": "登录成功", "IsAdmin": User.IsAdmin, "phone": User.Phone, "name": User.Name})
	return
}

// verifyAction
func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy + "解析失败")
	}
	clamis, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify access")
	return clamis, nil
}

// AuthRequired 验证token
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		strToken := c.Request.Header.Get("token")

		if strToken == "" {
			code = Status.InvalidToken
			c.JSON(code, Controller.R(code, nil, "无token"))
			c.Abort()
			return
		}
		claim, err := verifyAction(strToken)
		if err != nil {
			code = Status.InvalidToken
			c.JSON(code, Controller.R(code, nil, err.Error()))
			c.Abort()
			return
		}

		user := Service.GetUserByID(claim.UserID)
		c.Set("adminType", user.AdminType)
		c.Set("admin", user.IsAdmin)
		c.Set("studentid", user.StudentID)
		c.Set("userID", user.ID)
		c.Next()
	}
}

// GetUserByToken 通过token获取用户信息
func GetUserByToken(c *gin.Context) {
	var code int
	strToken := c.Request.Header.Get("token")

	if strToken == "" {
		code = Status.InvalidToken
		c.JSON(code, Controller.R(code, nil, "无效token"))
		return
	}
	claim, err := verifyAction(strToken)
	if err != nil {
		code = Status.InvalidToken
		c.JSON(code, Controller.R(code, nil, "token验证失败,检查token是否正确"))
		return
	}

	user := Service.GetUserByID(claim.UserID)
	code = Status.OK
	c.JSON(code, Controller.R(code, user, "获取数据成功"))
	return
}
