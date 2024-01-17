package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	cf "goStudy/config"
	"goStudy/utils/client"
	"goStudy/utils/logs"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"net/http"
)

type ClusterInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayname"`
	City        string `json:"city"`
	Area        string `json:"area"`
}

// 描述集群所用的配置信息
type ClusterConfig struct {
	ClusterInfo
	KubeConfig *string `json:"kubeconfig"`
}

// 描述集群状态
type ClusterStatus struct {
	ClusterInfo
	Version string
	Status  string
}

//var err error

// 结构体方法，判断集群状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo

	cf.ClientSet, err = client.ClientSetinitByString(c.KubeConfig)
	if err != nil {
		return clusterStatus, err
	}
	var cversion *version.Info
	cversion, err = cf.ClientSet.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	clusterStatus.Version = cversion.String()
	clusterStatus.Status = "Active"
	return clusterStatus, nil
}

func Add(c *gin.Context) {
	logs.Info(nil, "添加集群")
	var returnData cf.NewReturnData
	ustruct := ClusterConfig{}
	if err := c.ShouldBindJSON(&ustruct); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  500,
		},
		)
	} else {

		c.JSON(http.StatusOK, ustruct)
	}
	clusterStatus, err := ClusterConfig.getClusterStatus()
	if err != nil {
		msg := "无法获取集群信息" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {

	}

	//创建一个集群配置的scret
	var clusterSecretConfig corev1.Secret
	clusterSecretConfig.Name = ustruct.Id
	clusterSecretConfig.Labels = map[string]string{"metadata": "true"}
	clusterSecretConfig.Annotations = map[string]string{"displayname": ustruct.DisplayName, "city": ustruct.City, "area": ustruct.Area}
	//secret的data字段，需要加密，stringdata自带加密，所以此处直接使用stringdata
	clusterSecretConfig.StringData = map[string]string{"kubeconfig": *ustruct.KubeConfig}
	_, err = cf.ClientSet.CoreV1().Secrets("").Create(context.TODO(), &clusterSecretConfig, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"集群id:": ustruct.Id, "集群名称:": ustruct.DisplayName}, "集群创建失败")
		msg := "添加集群失败，secret创建失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {
		logs.Error(map[string]interface{}{"集群id:": ustruct.Id, "集群名称:": ustruct.DisplayName}, "集群创建成功")
		msg := "添加集群成功"
		returnData.Message = msg
		returnData.Status = 200
		c.JSON(http.StatusOK, returnData)
	}
}

func Update(c *gin.Context) {
	logs.Info(nil, "更新集群")
	ustruct := ClusterConfig{}
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

func Delete(c *gin.Context) {
	logs.Info(nil, "删除集群")
	ustruct := ClusterConfig{}
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

func Get(c *gin.Context) {
	logs.Info(nil, "获取集群信息")
	ustruct := ClusterConfig{}
	ustruct.DisplayName = c.Query("displayname")
	c.JSON(http.StatusOK, ustruct)

}
