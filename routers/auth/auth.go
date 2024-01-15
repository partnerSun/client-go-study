package auth

import (
	"github.com/gin-gonic/gin"
	"goStudy/controllers/auth"
)

func login(authgroup *gin.RouterGroup) {
	authgroup.POST("/login", auth.LoginFunc)
}

func logout(authgroup *gin.RouterGroup) {
	authgroup.GET("/logout", auth.LogoutFunc)
}

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	authGroup := sgroup.Group("/auth")
	login(authGroup)
	logout(authGroup)

}
