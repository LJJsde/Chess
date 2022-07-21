package controller

import (
	"Chess/dao"
	"Chess/module"
	"Chess/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Entrance() {
	router := InitRouter()
	router.POST("/register", Register)
	router.GET("/login", Login)
	router.Run(":8080")
}

func InitRouter() *gin.Engine {
	r := gin.New() //创建gin框架路由实例
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1/") //路由分组
	{
		apiv1Token := apiv1.Group("token/") //只有一个token组
		apiv1Token.Use(TokenVer())          //使用token鉴权中间件
		{

		}
	}
	return r
}

func Register(ctx *gin.Context) {
	var requestUser module.SUser
	ctx.Bind(&requestUser)
	UserID := util.SpawnOneRandomNumber(10000000, 99999999) + util.GetLocalDateTime()
	UserName := requestUser.UserName
	Email := requestUser.Email
	Password := requestUser.Password

	if len(Email) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "pls input Email",
		})
		return
	}
	if len(Password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}
	if len(Password) > 15 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    423,
			"message": "密码不能多于15位",
		})
	}
	var user module.SUser
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	//创建用户
	dao.CreateNewUser(UserID, UserName, Email, HashedPassword)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
	token, err := util.GeterateToken(user.ID, user.Email)
	if err != nil {
		module.ResponseWithJson(10001, "创建token失败", err, ctx)
		return
	}

	module.ResponseWithJson(200, gin.H{"User": user, "Token": token}, nil, ctx)
}

func Login(ctx *gin.Context) {
	var requestUser module.SUser
	ctx.Bind(&requestUser)
	LoginEmail := requestUser.Email
	LoginPassword := requestUser.Password
	if len(LoginPassword) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不能少于6位",
		})
		return
	}
	if len(LoginPassword) > 15 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    423,
			"message": "密码不能多于15位",
		})
	}
	if LoginEmail == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(LoginPassword), []byte(LoginPassword)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码错误",
		})
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}
