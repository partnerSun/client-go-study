package auth

import (
	"client-go-study/controllers/auth"
	"github.com/gin-gonic/gin"
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
