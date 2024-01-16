package pod

import (
	"github.com/gin-gonic/gin"
	cf "goStudy/config"
	"goStudy/pod"
	"goStudy/utils/logs"
	"net/http"
)

type podInfo struct {
	//Clustername string `json:"clusername"`
	NameSpace string `json:"namespace"`
	PodName   string `json:"podname"`
}

func Get(c *gin.Context) {
	ustruct := podInfo{}
	ustruct.NameSpace = c.Query("namespace")
	ustruct.PodName = c.Query("podname")
	podinfo, err := pod.GetPod(ustruct.NameSpace, ustruct.PodName)
	if err != nil {
		logs.Error(nil, err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "pod获取失败",
			"status":  404,
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": podinfo,
			"status":  200,
		})
	}

}

func Create(c *gin.Context) {
	ustruct := cf.PodSruct{}
	if err := c.ShouldBindJSON(&ustruct); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  500,
		},
		)
	} else {
		//c.JSON(http.StatusOK, ustruct)
		_, err := pod.CreatePod(ustruct)
		if err != nil {
			logs.Error(nil, err.Error())
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  500,
			})
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "pod创建成功",
				"status":  200,
			})
		}
	}

}
