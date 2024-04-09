package node

import (
	cf "client-go-study/config"
	"client-go-study/utils/logs"
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"time"
)

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
		//returnData.Data = nodes.Items
		// 打印节点信息
		//fmt.Printf("%-20s %-7s %-10s %-10s %-7s\n", "NAME", "STATUS", "ROLES", "AGE", "VERSION")
		var clusterlist []map[string]string
		for _, node := range nodes.Items {
			status := "Unknown"
			for _, condition := range node.Status.Conditions {
				if condition.Type == "Ready" && condition.Status == "True" {
					status = "Ready"
					break
				} else if condition.Type == "Ready" && condition.Status == "False" {
					status = "NotReady"
					break
				}
			}
			var roles string
			for _, role := range node.GetLabels()["kubernetes.io/role"] {
				roles += string(role) + ","
			}
			//fmt.Printf("roles:", roles)
			if roles != "" {
				roles = roles[:len(roles)-1] // remove trailing comma
			} else {
				roles = "null"
			}

			age := time.Since(node.CreationTimestamp.Time).Round(time.Second).String()
			version := node.Status.NodeInfo.KubeletVersion
			//fmt.Printf("%-20s %-7s %-10s %-10s %-7s\n", node.GetName(), status, roles, age, version)
			nodeList := map[string]string{
				"name":    node.GetName(),
				"status":  status,
				"roles":   roles,
				"age":     age,
				"version": version,
			}

			clusterlist = append(clusterlist, nodeList)
			//fmt.Printf("clusterlist:", clusterlist)
		}
		returnData.Data = make(map[string]interface{})
		returnData.Data["items"] = clusterlist
		//logs.Info(nil, "获取节点信息成功")
		//fmt.Printf("items: %v\n", returnData.Data["items"])
		c.JSON(http.StatusOK, returnData)
	}

}
