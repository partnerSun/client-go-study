// 路由层 管理程序的路由信息
package routers

import (
	"client-go-study/config"
	"client-go-study/routers/auth"
	"client-go-study/routers/cluster"
	"client-go-study/routers/namespace"
	"client-go-study/routers/node"
	"client-go-study/routers/pod"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	//r = gin.New()
	apiGroup := r.Group("/api")
	auth.RegisterSubRouters(apiGroup)
	//cluster.RegisterSubRouters(apiGroup)
	pod.RegisterSubRouters(apiGroup)
	cluster.RegisterSubRouters(apiGroup)
	namespace.RegisterSubRouters(apiGroup)
	node.RegisterSubRouters(apiGroup)
	r.Run(config.Port)
}
