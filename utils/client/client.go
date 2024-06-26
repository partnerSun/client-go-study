package client

import (
	"client-go-study/utils/logs"
	"fmt"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func ClientSetinit(kconfig *string) (*kubernetes.Clientset, error) {
	// 使用 clientcmd 加载 kubeconfig 文件
	config, err := clientcmd.BuildConfigFromFlags("", *kconfig)
	// 通过 InClusterConfig 方法获取集群内 kubeconfig 配置
	//kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		//fmt.Printf("Error building kubeconfig: %v\n", err)
		logs.Error(nil, "kubeconfig构建失败")
		return nil, err
	}

	//创建 Kubernetes 客户端配置
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		//fmt.Printf("Error creating clientset: %v\n", err)
		logs.Error(nil, "clientset创建失败")
		return nil, err
	}
	//返回 Kubernetes 客户端配置
	return clientset, nil

}

func DynamicClientInit(kconfig *string) *dynamic.DynamicClient {
	// 使用 clientcmd 加载 kubeconfig 文件
	config, err := clientcmd.BuildConfigFromFlags("", *kconfig)
	// 通过 InClusterConfig 方法获取集群内 kubeconfig 配置
	//kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// 创建 DynamicClient
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return dynamicClient
}

func ClientSetinitByString(kconfig string) (*kubernetes.Clientset, error) {
	// 使用 clientcmd 加载 kubeconfig 文件
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kconfig))
	// 通过 InClusterConfig 方法获取集群内 kubeconfig 配置
	//kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		return nil, err
	}

	//创建 Kubernetes 客户端配置
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating clientset: %v\n", err)
		return nil, err
	}
	//返回 Kubernetes 客户端配置
	return clientset, err

}
