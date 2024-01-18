package namespace

import (
	"github.com/gin-gonic/gin"
	namespace "goStudy/controllers/cluster"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	clusterGroup := sgroup.Group("/namespace")
	AddC(clusterGroup)
	DelC(clusterGroup)
	UpdateC(clusterGroup)
	GetC(clusterGroup)
	ListC(clusterGroup)

}

func AddC(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/add", namespace.Add)
}

func DelC(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/delete", namespace.Delete)
}

func UpdateC(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/update", namespace.Update)
}

func GetC(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", namespace.Get)
}

func ListC(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", namespace.List)
}
