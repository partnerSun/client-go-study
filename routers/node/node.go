package node

import (
	"client-go-study/controllers/node"
	"github.com/gin-gonic/gin"
)

// 子路由
func RegisterSubRouters(sgroup *gin.RouterGroup) {
	nodeGroup := sgroup.Group("/node")
	nodeList(nodeGroup)

}

func nodeList(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/list", node.List)
}
