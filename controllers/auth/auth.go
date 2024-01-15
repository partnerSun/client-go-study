package auth

import (
	"github.com/gin-gonic/gin"
	"goStudy/config"
	"goStudy/utils/jwtutils"
	"goStudy/utils/logs"
	"net/http"
)

type userinfo struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	Password string `json:"password"`
}

func LoginFunc(c *gin.Context) {
	ustruct := userinfo{}
	if err := c.ShouldBindJSON(&ustruct); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  401,
		})
	} else {

		logs.Debug(map[string]interface{}{"用户名": ustruct.Username, "密码": ustruct.Password}, "用户信息绑定成功")
		//验证用户名、密码是否正确
		//1. 存放在数据库中
		//2. 在环境变量中
		if ustruct.Username == config.Username && ustruct.Password == config.Password {
			//登录成功，然后生成jwt的token
			ss, err := jwtutils.Gentoken(ustruct.Username)
			if err != nil {
				logs.Error(map[string]interface{}{"用户名": ustruct.Username}, "用户已登录但是token生成失败")
				c.JSON(http.StatusOK, gin.H{
					"status":  401,
					"message": "登陆成功但是token生成失败",
				})
				return
			}

			logs.Info(map[string]interface{}{"用户名": ustruct.Username}, "用户认证成功")
			data := make(map[string]interface{})
			data["token"] = ss
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "登录成功",
				"data":    data,
			})
			return

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "用户认证失败，用户名或密码错误",
				"status":  401,
			})
		}
	}

}

func LogoutFunc(c *gin.Context) {
	uname := c.Query("username")

	if uname == config.Username {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户已退出",
			"status":  200,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户不存在",
			"status":  401,
		})
	}
	logs.Debug(nil, "用户已退出")
}
