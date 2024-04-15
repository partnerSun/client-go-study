package node

import (
	cf "client-go-study/config"
	"client-go-study/utils/logs"
	"context"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type nodeStruct struct {
	ClusterId string            `json:"clusterId"`
	NodeName  string            `json:"name"`
	Labels    map[string]string `json:"labels"`
	Taints    []corev1.Taint    `json:"taints"`
}

func List(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "开始获取节点信息")
	// 查询集群中的节点信息
	nodes, err := cf.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(nil, "获取节点信息失败")
		msg := "获取节点信息失败:" + err.Error()
		returnData.Message = msg
		returnData.Type = "error"
		returnData.Status = 400
	} else {
		msg := "获取节点信息成功"
		returnData.Message = msg
		returnData.Status = 200
		returnData.Type = "success"
		returnData.Data = make(map[string]interface{})
		returnData.Data["items"] = nodes.Items

	}
	c.JSON(http.StatusOK, returnData)

}

func Update(c *gin.Context) {
	var returnData cf.NewReturnData

	nodeInfo := nodeStruct{}
	if err := c.ShouldBindJSON(&nodeInfo); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		msg := "集群配置信息不完整: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		returnData.Type = "error"
		c.JSON(200, returnData)
		return
	}
	// 解析节点信息的 JSON 字符串
	node, err := GetNode(nodeInfo.NodeName)
	if err != nil {
		msg := "Node获取失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		returnData.Type = "error"
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "无法获取Node信息")
		return
	}
	for key, value := range nodeInfo.Labels {
		node.Labels[key] = value
	}
	node.Spec.Taints = nodeInfo.Taints
	_, err = cf.ClientSet.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		msg := "Node更新失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		returnData.Type = "error"
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "Node更新失败")
		return
	}
	msg := nodeInfo.NodeName + "节点更新成功"
	returnData.Message = msg
	returnData.Status = 200
	returnData.Type = "success"
	//returnData.Data["items"]=updatenode.
	c.JSON(http.StatusOK, returnData)

}

func GetNode(name string) (*corev1.Node, error) {
	node, err := cf.ClientSet.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return node, nil
}
func Get(c *gin.Context) {
	var returnData cf.NewReturnData
	nodeinfo := nodeStruct{}
	nodeinfo.NodeName = c.Query("name")
	node, err := GetNode(nodeinfo.NodeName)
	if err != nil {
		msg := "Node获取失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		returnData.Type = "error"
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "无法获取Node信息")
		return
	}
	msg := "获取Node " + nodeinfo.NodeName + " 成功"
	returnData.Message = msg
	returnData.Status = 200
	returnData.Type = "success"
	returnData.Data = make(map[string]interface{})

	returnData.Data["items"] = node
	c.JSON(http.StatusOK, returnData)

}
