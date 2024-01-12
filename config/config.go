package config

import (
	"github.com/spf13/viper"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

var (
	Namesapce            string
	MetaNamespace        string
	InClusterKubeconfig  *string
	OutClusterKubeconfig *string
	DeploymentName       string
	Replicas             int32
	Image                string
	ClientSet            *kubernetes.Clientset
	DynamicClient        *dynamic.DynamicClient
)

func init() {
	//设置命名空间
	viper.SetDefault("name_space", "default") //为空则是所有命名空间
	Namesapce = viper.GetString("name_space")
	//要获取、修改的deployment名字
	viper.SetDefault("deployment_name", "nginx-example-deployment")
	DeploymentName = viper.GetString("deployment_name")

	//设置deployment副本数
	viper.SetDefault("replicas", 3)
	Replicas = viper.GetInt32("replicas")
	//设置image
	viper.SetDefault("image_nameV", "nginx")
	Image = viper.GetString("image_nameV")
	//定义元数据命名空间
	viper.SetDefault("metans", "meta-namespace")
	MetaNamespace = viper.GetString("metans")
}
