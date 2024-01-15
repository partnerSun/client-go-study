// 路由层 管理程序的路由信息
package routers

import (
	"github.com/gin-gonic/gin"
	"goStudy/config"
	"goStudy/routers/auth"
)

func RegisterRouters(r *gin.Engine) {
	//r = gin.New()
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
	r.Run(config.Port)
}
