package pod

import (
	"github.com/gin-gonic/gin"

	"client-go-study/controllers/pod"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	podGroup := sgroup.Group("/pod")
	podget(podGroup)
	podlist(podGroup)
	podcreate(podGroup)
	poddelete(podGroup)

}

func podget(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", pod.Get)
}

func podcreate(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/create", pod.Create)
}

func poddelete(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/delete", pod.Delete)
}

func podlist(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", pod.List)
}
