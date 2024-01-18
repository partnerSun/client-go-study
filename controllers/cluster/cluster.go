package cluster

import (
	"context"
	"github.com/gin-gonic/gin"
	cf "goStudy/config"
	"goStudy/utils"
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
	KubeConfig string `json:"kubeconfig"`
}

// 描述集群状态
type ClusterStatus struct {
	ClusterInfo
	Version string
	Status  string
}

var err error

// 结构体方法，判断集群状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo

	cf.ClientSet, err = client.ClientSetinitByString(c.KubeConfig)

	if err != nil {
		logs.Error(nil, "kubeconfig配置有问题，生成Clientset出错")
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

// 通过secret创建集群，secret中保存集群信息和kubeconfig
func Add(c *gin.Context) {
	logs.Info(nil, "添加	集群")
	var returnData cf.NewReturnData
	clusterconfig := ClusterConfig{}
	if err := c.ShouldBindJSON(&clusterconfig); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		msg := "集群的配置信息不完整: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	clusterStatus, err := clusterconfig.getClusterStatus()
	if err != nil {
		msg := "无法获取集群信息" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "集群失败,无法获取集群信息")
		return
	}
	logs.Info(map[string]interface{}{"集群名称": clusterconfig.DisplayName, "集群ID": clusterconfig.Id}, "开始配置集群")
	//创建一个集群配置的scret
	var clusterSecretConfig corev1.Secret
	clusterSecretConfig.Name = clusterconfig.Id
	clusterSecretConfig.Labels = map[string]string{"metadata": "true"}
	//添加注释，写法由固定map改为动态map
	//clusterSecretConfig.Annotations = map[string]string{"displayname": ustruct.DisplayName, "city": ustruct.City, "area": ustruct.Area}
	clusterSecretConfig.Annotations = make(map[string]string)
	m := utils.Struct2map(clusterStatus)
	clusterSecretConfig.Annotations = m

	//secret的data字段，需要加密，stringdata自带加密，所以此处直接使用stringdata
	clusterSecretConfig.StringData = map[string]string{"kubeconfig": clusterconfig.KubeConfig}
	_, err = cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Create(context.TODO(), &clusterSecretConfig, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"集群id:": clusterconfig.Id, "集群名称:": clusterconfig.DisplayName}, "集群创建失败")
		msg := "添加集群失败，secret创建失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {
		logs.Error(map[string]interface{}{"集群id:": clusterconfig.Id, "集群名称:": clusterconfig.DisplayName}, "集群创建成功")
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

// 通过删除secret来删除集群
func Delete(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "删除集群")
	clusterid := c.Query("clusterid")
	cf.ClientSet, _ = client.ClientSetinit(cf.OutClusterKubeconfig)
	err = cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Delete(context.TODO(), clusterid, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(nil, "集群(secret)删除失败")
		msg := "集群删除失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		msg := "集群删除成功"
		returnData.Message = msg
		returnData.Status = 200
	}
	c.JSON(http.StatusOK, returnData)
}

func Get(c *gin.Context) {
	logs.Info(nil, "获取集群信息")
	ustruct := ClusterConfig{}
	ustruct.DisplayName = c.Query("displayname")
	c.JSON(http.StatusOK, ustruct)

}
