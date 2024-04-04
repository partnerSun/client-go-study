package namespace

import (
	cf "client-go-study/config"
	"client-go-study/utils/logs"
	"context"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type NsBasicInfo struct {
	ClusterId string      `json:"clusterid" form:"clusterid"`
	NameSpace string      `json:"namespace" form:"namespace"`
	Name      string      `json:"name" form:"name"`
	Items     interface{} `json:"items" form:"items"`
}

//type NsInfo struct {
//	NsBasicInfo
//	status string
//}

var err error

func add(c *gin.Context) {
	var returnData cf.NewReturnData
	//绑定post参数
	clusterconfig := NsBasicInfo{}
	if err := c.ShouldBindJSON(&clusterconfig); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		msg := "Namespace的配置信息不完整: " + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}

	////用于创建ns
	var clusterNamespace corev1.Namespace
	clusterNamespace.Name = clusterconfig.Name

	_, err = cf.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &clusterNamespace, metav1.CreateOptions{})

	if err != nil {
		logs.Error(map[string]interface{}{"Namespace名称:": clusterconfig.Name}, "Namespace创建失败")
		msg := "Namespace创建失败：" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {
		logs.Info(map[string]interface{}{"Namespace名称:": clusterconfig.Name}, "Namespace创建成功")
		msg := "Namespace创建成功"
		returnData.Message = msg
		returnData.Status = 200
		c.JSON(http.StatusOK, returnData)
	}
}

// 创建Namespace
func Create(c *gin.Context) {
	logs.Info(nil, "添加Namespace")
	add(c)
}

func abc(c *gin.Context, item interface{}) (nsbasic NsBasicInfo, err2 error) {
	//绑定post参数
	clusterconfig := NsBasicInfo{}
	clusterconfig.Items = item
	if err := c.ShouldBindJSON(&clusterconfig); err != nil { //如果 JSON 数据无法绑定到结构体，它不会返回错误，而是返回一个布尔值（bool）
		return nsbasic, err
	}
	return clusterconfig, nil
}

// 更新Namespace
func Update(c *gin.Context) {
	logs.Info(nil, "更新Namespace")
	//addOrUpdate(c, "update")
	var ns corev1.Namespace
	var returnData cf.NewReturnData
	_, err = abc(c, &ns)
	if err != nil {
		msg := "Namespace配置信息不完整：" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(http.StatusOK, returnData)
		return
	}
	_, err = cf.ClientSet.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"Namespace名称:": ns.Name}, "Namespace更新失败")
		msg := "Namespace更新失败：" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
	} else {
		logs.Info(map[string]interface{}{"Namespace名称:": ns.Name}, "Namespace更新成功")
		msg := "Namespace更新成功"
		returnData.Message = msg
		returnData.Status = 200
		c.JSON(http.StatusOK, returnData)
	}

}

// 删除Namespace
func Delete(c *gin.Context) {
	var returnData cf.NewReturnData
	clusterconfig := NsBasicInfo{}
	clusterconfig.Name = c.Query("name")
	err := cf.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), clusterconfig.Name, metav1.DeleteOptions{})
	if err != nil {
		msg := "Namespace删除失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "Namespace删除失败")
		return
	}
	logs.Info(nil, "成功删除Namespace"+clusterconfig.Name)
	msg := "删除Namespace " + clusterconfig.Name + " 成功"
	returnData.Message = msg
	returnData.Status = 200
	c.JSON(http.StatusOK, returnData)
}

// 查询Namespace信息
func Get(c *gin.Context) {
	var returnData cf.NewReturnData
	clusterconfig := NsBasicInfo{}
	clusterconfig.Name = c.Query("name")
	nsinfo, err := cf.ClientSet.CoreV1().Namespaces().Get(context.TODO(), clusterconfig.Name, metav1.GetOptions{})
	//nsinfo, err := getNameSpaceStatus(clusterconfig.Name)
	if err != nil {
		msg := "Namespace获取失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "Namespace失败,无法获取Namespace信息")
		return
	}
	msg := "获取Namespace " + clusterconfig.Name + " 成功"
	returnData.Message = msg
	returnData.Status = 200
	returnData.Data = make(map[string]interface{})

	returnData.Data["items"] = nsinfo
	c.JSON(http.StatusOK, returnData)
}

// 获取Namespace列表
func List(c *gin.Context) {
	var returnData cf.NewReturnData

	nsAllList, err := cf.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		msg := "获取Namespace列表失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(http.StatusOK, returnData)
		logs.Error(map[string]interface{}{"error": err.Error()}, "获取Namespace列表失败")
		return
	}
	returnData.Data = make(map[string]interface{}) //map

	logs.Info(nil, "获取Namespace列表成功")
	msg := "获取Namespace列表成功"
	returnData.Message = msg
	returnData.Status = 200
	returnData.Data["items"] = nsAllList
	c.JSON(http.StatusOK, returnData)
}
