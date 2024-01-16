package pod

import (
	"github.com/gin-gonic/gin"

	"goStudy/controllers/pod"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	podGroup := sgroup.Group("/pod")
	podlist(podGroup)
	podcreate(podGroup)

}

func podlist(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", pod.Get)
}

func podcreate(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/create", pod.Create)
}
