// 路由层 管理程序的路由信息
package routers

import (
	"github.com/gin-gonic/gin"
	"goStudy/config"
	"goStudy/routers/auth"
	"goStudy/routers/cluster"
	"goStudy/routers/namespace"
	"goStudy/routers/pod"
)

func RegisterRouters(r *gin.Engine) {
	//r = gin.New()
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
	//cluster.RegisterSubRouters(apiGroup)
	pod.RegisterSubRouters(apiGroup)
	cluster.RegisterSubRouters(apiGroup)
	namespace.RegisterSubRouters(apiGroup)
	r.Run(config.Port)
}
