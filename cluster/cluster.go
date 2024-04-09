package cluster

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func ClusterInfoCheck(clientset *kubernetes.Clientset) {

	// 查询集群中的节点信息
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error! failed to Getting nodes: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Nodes in the cluster:")
	for _, node := range nodes.Items {
		fmt.Printf("- Node: %s\n", node.GetName())

	}

}
