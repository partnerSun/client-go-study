package cluster

import (
	cf "client-go-study/config"
	"client-go-study/utils"
	"client-go-study/utils/client"
	"client-go-study/utils/logs"
	"context"
	"github.com/gin-gonic/gin"
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

//var err error

// 结构体方法，判断集群状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	clientSet, err := client.ClientSetinitByString(c.KubeConfig)

	if err != nil {
		logs.Error(nil, "kubeconfig配置有问题，生成Clientset出错")
		return clusterStatus, err
	}
	var cversion *version.Info
	cversion, err = clientSet.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	clusterStatus.Version = cversion.String()
	clusterStatus.Status = "Active"
	return clusterStatus, nil
}

// 提取相同逻辑，添加和更新集群功能
func addOrUpdate(c *gin.Context, op string) {
	var returnData cf.NewReturnData
	//区分创建还是更新
	var arg string
	var err error
	if op == "Create" || op == "create" {
		arg = "创建"
	} else {
		arg = "更新"
	}
	//绑定post参数
	clusterconfig := ClusterConfig{}
	if err := c.ShouldBindJSON(&clusterconfig); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		msg := arg + "集群的配置信息不完整: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	//判断集群状态
	clusterStatus, _ := clusterconfig.getClusterStatus()
	//if err != nil {
	//	msg := "无法获取集群信息" + err.Error()
	//	returnData.Message = msg
	//	returnData.Status = 400
	//	c.JSON(http.StatusOK, returnData)
	//	logs.Info(map[string]interface{}{"error": err.Error()}, "集群失败,无法获取集群信息")
	//	return
	//}
	logs.Info(map[string]interface{}{"集群名称": clusterconfig.DisplayName, "集群ID": clusterconfig.Id}, "开始"+arg+"集群")

	//配置scret
	var clusterSecretConfig corev1.Secret
	clusterSecretConfig.Name = clusterconfig.Id
	clusterSecretConfig.Labels = map[string]string{"metadata": "true"}

	//添加Annotations
	clusterSecretConfig.Annotations = make(map[string]string)
	m := utils.Struct2map(clusterStatus) //结构体转map
	clusterSecretConfig.Annotations = m

	//secret的data字段，需要加密，stringdata自带加密，所以此处直接使用stringdata
	clusterSecretConfig.StringData = map[string]string{"kubeconfig": clusterconfig.KubeConfig}

	if op == "Create" || op == "create" {
		_, err = cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Create(context.TODO(), &clusterSecretConfig, metav1.CreateOptions{})
	} else {
		_, err = cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Update(context.TODO(), &clusterSecretConfig, metav1.UpdateOptions{})
	}

	if err != nil {
		logs.Error(map[string]interface{}{"集群id:": clusterconfig.Id, "集群名称:": clusterconfig.DisplayName}, "集群"+arg+"失败")
		msg := arg + "集群失败：" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {
		logs.Error(map[string]interface{}{"集群id:": clusterconfig.Id, "集群名称:": clusterconfig.DisplayName}, "集群"+arg+"成功")
		msg := arg + "集群成功"
		returnData.Message = msg
		returnData.Status = 200
		c.JSON(http.StatusOK, returnData)
	}
}

// 通过secret创建集群，secret中保存集群信息和kubeconfig
func Add(c *gin.Context) {
	logs.Info(nil, "添加集群")
	addOrUpdate(c, "create")
}

func Update(c *gin.Context) {
	logs.Info(nil, "更新集群")
	addOrUpdate(c, "update")
}

// 通过删除secret来删除集群
func Delete(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "删除集群")
	clusterid := c.Query("clusterid")
	err := cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Delete(context.TODO(), clusterid, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(nil, "集群(secret)删除失败")
		msg := "集群删除失败：" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		msg := "集群删除成功"
		returnData.Message = msg
		returnData.Status = 200
	}
	c.JSON(http.StatusOK, returnData)
}

// 查询集群信息
func Get(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "开始获取某个集群信息")
	clusterid := c.Query("clusterid")
	sinfo, err := cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).Get(context.TODO(), clusterid, metav1.GetOptions{})
	if err != nil {
		logs.Error(nil, "获取集群信息失败")
		msg := "获取集群信息失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		var clusterinfo map[string]string                            //定义一个字典保存集群信息
		clusterinfo = sinfo.Annotations                              //先保存Annotations
		clusterinfo["kubeconfig"] = string(sinfo.Data["kubeconfig"]) //再把kubeconfig添加到clusterinfo中

		msg := "获取集群信息成功"
		returnData.Message = msg
		returnData.Status = 200
		returnData.Data = map[string]interface{}{"iterm": clusterinfo}
		logs.Info(nil, "获取集群信息成功")

	}
	c.JSON(http.StatusOK, returnData)
}

// 获取集群信息 就是固定ns下的secret列表
func List(c *gin.Context) {
	var returnData cf.NewReturnData
	logs.Info(nil, "获取集群列表")
	listoptions := metav1.ListOptions{LabelSelector: "metadata=true"} //通过labelselector过滤指定secret
	slist, err := cf.ClientSet.CoreV1().Secrets(cf.MetaNamespace).List(context.TODO(), listoptions)
	if err != nil {
		logs.Error(nil, "获取集群列表失败")
		msg := "获取集群列表失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	}
	logs.Info(nil, "获取集群列表成功")
	msg := "获取集群列表成功"
	returnData.Message = msg
	returnData.Status = 200
	returnData.Data = make(map[string]interface{}) //map

	var clusterlist []map[string]string //定义一个变量 用于接收列表信息
	for _, v := range slist.Items {
		anno := v.Annotations
		clusterlist = append(clusterlist, anno)
	}
	returnData.Data["items"] = clusterlist
	c.JSON(http.StatusOK, returnData)
	//logs.Info(nil, slist)
}
