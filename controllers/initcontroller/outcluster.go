package initcontroller

import (
	cf "goStudy/config"
	"goStudy/utils/client"
)

func metaDataInit2(MetaNamespace string) {
	tmpstr := "config/kubeconfig"
	cf.OutClusterKubeconfig = &tmpstr
	//生成 Kubernetes 客户端配置，KubeconfigOutCluster或KubeconfigInCluster
	cf.ClientSet = client.ClientSetinit(cf.OutClusterKubeconfig)
	cf.DynamicClient = client.DynamicClientInit(cf.OutClusterKubeconfig)
	//判断命名空间是否存在
	//_, err := cf.ClientSet.CoreV1().Namespaces().Get(context.TODO(), MetaNamespace, metav1.GetOptions{})
	//inClusterVersion, err := cf.ClientSet.Discovery().ServerVersion()
	//if err != nil {
	//	fmt.Printf("%s元数据命名空间未创建\n", MetaNamespace)
	//	var tmpMetaNamespace corev1.Namespace
	//	tmpMetaNamespace.Name = MetaNamespace
	//	_, err = cf.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &tmpMetaNamespace, metav1.CreateOptions{})
	//	if err != nil {
	//		fmt.Printf("%s 命名空间创建失败\n", MetaNamespace)
	//		panic(err.Error())
	//	}
	//
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Printf("%s 命名空间创建成功\n", MetaNamespace)
	//	fmt.Printf("集群版本是%s\n", inClusterVersion)
	//} else {
	//	fmt.Printf("%s 命名空间已存在\n", MetaNamespace)
	//	fmt.Printf("集群版本是%s\n", inClusterVersion)
	//}

}
