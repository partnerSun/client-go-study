package cluster

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClusterInfo struct {
	//Clustername string `json:"clusername"`
	NameSpace string `json:"namespace"`
}

func Add(c *gin.Context) {
	ustruct := ClusterInfo{}
	if err := c.ShouldBindJSON(&ustruct); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  500,
		},
		)
	} else {

		c.JSON(http.StatusOK, ustruct)
	}
}
