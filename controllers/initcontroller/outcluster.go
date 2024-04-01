package initcontroller

import (
	cf "client-go-study/config"
	"client-go-study/utils/client"
	"client-go-study/utils/logs"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func metaDataInit2(MetaNamespace string) {
	tmpstr := "config/kubeconfig"
	cf.OutClusterKubeconfig = &tmpstr
	//生成 Kubernetes 客户端配置，KubeconfigOutCluster或KubeconfigInCluster
	cf.ClientSet, _ = client.ClientSetinit(cf.OutClusterKubeconfig)
	cf.DynamicClient = client.DynamicClientInit(cf.OutClusterKubeconfig)
	//判断命名空间是否存在
	_, err := cf.ClientSet.CoreV1().Namespaces().Get(context.Background(), cf.MetaNamespace, metav1.GetOptions{})
	if err != nil {
		logs.Error(nil, "元数据命名空间未创建")
		var tmpMetaNamespace corev1.Namespace
		tmpMetaNamespace.Name = cf.MetaNamespace
		_, err = cf.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &tmpMetaNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(nil, "元数据命名空间创建失败")
			panic(err.Error())
		}
		inClusterVersion, err := cf.ClientSet.Discovery().ServerVersion()
		if err != nil {
			panic(err.Error())
		}
		logs.Info(map[string]interface{}{"当前集群版本": inClusterVersion}, "元数据命名空间创建成功")

	}

}
