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

func podget(podGroup *gin.RouterGroup) {
	podGroup.GET("/get", pod.Get)
}

func podcreate(podGroup *gin.RouterGroup) {
	podGroup.POST("/create", pod.Create)
}

func poddelete(podGroup *gin.RouterGroup) {
	podGroup.POST("/delete", pod.Delete)
}

func podlist(podGroup *gin.RouterGroup) {
	podGroup.GET("/list", pod.List)
}
