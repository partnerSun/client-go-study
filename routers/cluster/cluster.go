package cluster

import (
	"github.com/gin-gonic/gin"
	"goStudy/controllers/cluster"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	clusterGroup := sgroup.Group("/cluster")
	AddCluster(clusterGroup)

}

func AddCluster(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/addcluster", cluster.Add)
}
