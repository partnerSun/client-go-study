package initcontroller

import cf "goStudy/config"

func metaDataInit2() {
	tmpstr := "../../config/config"
	cf.OutClusterKubeconfig = &tmpstr
	//生成 Kubernetes 客户端配置，KubeconfigOutCluster或KubeconfigInCluster

}
