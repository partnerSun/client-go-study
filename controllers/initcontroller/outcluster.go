package initcontroller

import (
	"context"
	"fmt"
	cf "goStudy/config"
	"goStudy/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func metaDataInit2() {
	tmpstr := "./config/config"
	cf.OutClusterKubeconfig = &tmpstr
	//生成 Kubernetes 客户端配置，KubeconfigOutCluster或KubeconfigInCluster
	cf.ClientSet = utils.ClientSetinit(cf.OutClusterKubeconfig)
	cf.DynamicClient = utils.DynamicClientInit(cf.OutClusterKubeconfig)
	//判断命名空间是否存在
	_, err := cf.ClientSet.CoreV1().Namespaces().Get(context.TODO(), cf.MetaNamespace, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("%s元数据命名空间未创建\n", cf.MetaNamespace)
		var tmpMetaNamespace corev1.Namespace
		tmpMetaNamespace.Name = cf.MetaNamespace
		_, err = cf.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &tmpMetaNamespace, metav1.CreateOptions{})
		if err != nil {
			fmt.Printf("%s 命名空间创建失败\n", cf.MetaNamespace)
			panic(err.Error())
		}
		inClusterVersion, err := cf.ClientSet.Discovery().ServerVersion()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("集群版本是%s\n", inClusterVersion)
		fmt.Printf("%s 命名空间创建成功\n", cf.MetaNamespace)
	}
}
