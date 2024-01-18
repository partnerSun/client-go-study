package cluster

import (
	"github.com/gin-gonic/gin"
	"goStudy/controllers/cluster"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	clusterGroup := sgroup.Group("/cluster")
	AddC(clusterGroup)
	DelC(clusterGroup)
	UpdateC(clusterGroup)
	GetC(clusterGroup)

}

func AddC(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/add", cluster.Add)
}

func DelC(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/delete", cluster.Delete)
}

func UpdateC(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/update", cluster.Update)
}

func GetC(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", cluster.Get)
}
