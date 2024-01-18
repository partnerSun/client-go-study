package namespace

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

type NameSpaceInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayname"`
	City        string `json:"city"`
	Area        string `json:"area"`
}

// 描述Namespace所用的配置信息
type NameSpaceConfig struct {
	NameSpaceInfo
	KubeConfig string `json:"kubeconfig"`
}

// 描述Namespace状态
type NameSpaceStatus struct {
	NameSpaceInfo
	Version string
	Status  string
}

var err error

// 结构体方法，判断Namespace状态
func (c *NameSpaceConfig) getClusterStatus() (NameSpaceStatus, error) {
	clusterStatus := NameSpaceStatus{}
	clusterStatus.NameSpaceInfo = c.NameSpaceInfo

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

// 提取相同逻辑，添加和更新Namespace功能
func addOrUpdate(c *gin.Context, op string) {
	var returnData cf.NewReturnData
	//区分创建还是更新
	var arg string
	if op == "Create" || op == "create" {
		arg = "创建"
	} else {
		arg = "更新"
	}
	//绑定post参数
	clusterconfig := NameSpaceConfig{}
	if err := c.ShouldBindJSON(&clusterconfig); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		msg := arg + "Namespace的配置信息不完整: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	//判断Namespace状态
	clusterStatus, err := clusterconfig.getClusterStatus()
	if err != nil {
		msg := "无法获取Namespace信息" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "Namespace失败,无法获取Namespace信息")
		return
	}
	logs.Info(map[string]interface{}{"Namespace名称": clusterconfig.DisplayName, "NamespaceID": clusterconfig.Id}, "开始"+arg+"Namespace")

	//配置scret
	var clusterNamespace corev1.Namespace

	clusterNamespace.Name = clusterconfig.Id
	clusterNamespace.Labels = map[string]string{"metadata": "true"}

	//添加Annotations
	clusterNamespace.Annotations = make(map[string]string)
	m := utils.Struct2map(clusterStatus) //结构体转map
	clusterNamespace.Annotations = m

	//secret的data字段，需要加密，stringdata自带加密，所以此处直接使用stringdata
	//clusterNamespace.StringData = map[string]string{"kubeconfig": clusterconfig.KubeConfig}

	if op == "Create" || op == "create" {
		//_, err = cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Create(context.TODO(), &clusterSecretConfig, metav1.CreateOptions{})
		_, err = cf.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &clusterNamespace, metav1.CreateOptions{})

	} else {
		_, err = cf.ClientSet.CoreV1().Namespaces().Update(context.TODO(), &clusterNamespace, metav1.UpdateOptions{})
	}

	if err != nil {
		logs.Error(map[string]interface{}{"Namespaceid:": clusterconfig.Id, "Namespace名称:": clusterconfig.DisplayName}, "Namespace"+arg+"失败")
		msg := arg + "Namespace失败：" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {
		logs.Error(map[string]interface{}{"Namespaceid:": clusterconfig.Id, "Namespace名称:": clusterconfig.DisplayName}, "Namespace"+arg+"成功")
		msg := arg + "Namespace成功"
		returnData.Message = msg
		returnData.Status = 200
		c.JSON(http.StatusOK, returnData)
	}
}

// 创建Namespace
func Create(c *gin.Context) {
	logs.Info(nil, "添加Namespace")
	addOrUpdate(c, "create")
}

// 更新Namespace
func Update(c *gin.Context) {
	logs.Info(nil, "更新Namespace")
	addOrUpdate(c, "update")
}

// 删除Namespace
func Delete(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "删除Namespace")
	c.JSON(http.StatusOK, returnData)
}

// 查询Namespace信息
func Get(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "获取Namespace信息")

	c.JSON(http.StatusOK, returnData)
}

// 获取Namespace列表
func List(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "获取Namespace列表")
	c.JSON(http.StatusOK, returnData)
	//logs.Info(nil, slist)
}
