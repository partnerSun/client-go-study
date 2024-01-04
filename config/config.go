package config

import (
	"flag"
	"github.com/spf13/viper"
	"goStudy/utils"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"
)

var (
	Namesapce            string
	KubeconfigInCluster  *string
	KubeconfigOutCluster *string
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
	//在集群内使用时设置kubeconfig文件位置
	if home := homedir.HomeDir(); home != "" {
		KubeconfigInCluster = flag.String("kubeconfig", home+"/.kube/config", "(optional) absolute path to the kubeconfig file")
	} else {
		KubeconfigInCluster = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	//集群外可使用离线的config文件
	tmpstr := "./config/config"
	KubeconfigOutCluster = &tmpstr
	//生成 Kubernetes 客户端配置，KubeconfigOutCluster或KubeconfigInCluster
	ClientSet = utils.ClientSetinit(KubeconfigOutCluster)
	DynamicClient = utils.DynamicClientInit(KubeconfigOutCluster)
	//

}
