package initcontroller

import (
	"context"
	"flag"
	"fmt"
	cf "goStudy/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/homedir"
)

func metaDatainit() {
	if home := homedir.HomeDir(); home != "" {
		cf.InClusterKubeconfig = flag.String("kubeconfig", home+"/.kube/config", "(optional) absolute path to the kubeconfig file")
	} else {
		cf.InClusterKubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	//集群外可使用离线的config文件

	//判断命名空间是否存在
	_, err := clientset.CoreV1().Namespaces().Get(context.Background(), cf.MetaNamespace, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("%s元数据命名空间未创建", cf.MetaNamespace)
		var tmpMetaNamespace corev1.Namespace
		tmpMetaNamespace.Namespace = cf.MetaNamespace
		_, err = cf.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &tmpMetaNamespace, metav1.CreateOptions{})
		if err != nil {
			fmt.Printf("%s 命名空间创建失败", cf.MetaNamespace)
			panic(err.Error())
		}
		inClusterVersion, err := cf.ClientSet.Discovery().ServerVersion()
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("集群版本是%s", inClusterVersion)
		fmt.Printf("%s 命名空间创建成功", cf.MetaNamespace)
	}
}