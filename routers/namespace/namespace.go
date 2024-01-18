package namespace

import (
	"github.com/gin-gonic/gin"
	"goStudy/controllers/namespace"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	nsGroup := sgroup.Group("/namespace")
	AddC(nsGroup)
	DelC(nsGroup)
	UpdateC(nsGroup)
	GetC(nsGroup)
	ListC(nsGroup)

}

func AddC(nsGroup *gin.RouterGroup) {
	nsGroup.POST("/create", namespace.Create)
}

func DelC(nsGroup *gin.RouterGroup) {
	nsGroup.GET("/delete", namespace.Delete)
}

func UpdateC(nsGroup *gin.RouterGroup) {
	nsGroup.POST("/update", namespace.Update)
}

func GetC(nsGroup *gin.RouterGroup) {
	nsGroup.GET("/get", namespace.Get)
}

func ListC(nsGroup *gin.RouterGroup) {
	nsGroup.GET("/list", namespace.List)
}
