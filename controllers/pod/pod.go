package pod

import (
	cf "client-go-study/config"
	"client-go-study/pod"
	"client-go-study/utils/logs"
	"github.com/gin-gonic/gin"
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

func Delete(c *gin.Context) {
	ustruct := podInfo{}
	ustruct.NameSpace = c.Query("namespace")
	ustruct.PodName = c.Query("podname")
	if err := c.ShouldBindJSON(&ustruct); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  500,
		},
		)
	} else {
		//c.JSON(http.StatusOK, ustruct)
		_, err := pod.DelPod(ustruct.NameSpace, ustruct.PodName)
		if err != nil {
			logs.Error(nil, err.Error())
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  500,
			})
			c.Abort()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "pod删除成功",

				"status": 200,
			})
		}
	}

}

func List(c *gin.Context) {
	ustruct := podInfo{}
	var returnData cf.NewReturnData
	//var podStatusMap map[string]string
	msg := "获取集群列表成功"
	returnData.Message = msg
	returnData.Status = 200
	returnData.Data = make(map[string]interface{})
	var podStatusSlice []map[string]interface{}
	podlistmap := make(map[string]interface{})
	ustruct.NameSpace = c.Query("namespace")
	podlist, err := pod.AllPodStatusList(ustruct.NameSpace)
	//fmt.Println("podlist", podlist)
	if err != nil {
		logs.Error(nil, err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "pod获取失败",
			"status":  404,
		})
		c.Abort()
	} else {
		for _, pod := range podlist.Items {
			//fmt.Printf(" - Name: %s, State: %s\n", pod.Name, pod.Status.Phase)
			podlistmap["name"] = pod.Name
			podlistmap["status"] = string(pod.Status.Phase)
			//fmt.Println("podlistmap:", podlistmap)
			podStatusSlice = append(podStatusSlice, podlistmap)
		}
		returnData.Data["iterms"] = podStatusSlice
		c.JSON(http.StatusOK, returnData)
	}

}
