package deployment

import (
	cf "client-go-study/config"
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"os"
)

func CreateDeployByJson(dynamicClient *dynamic.DynamicClient) {
	// 你的 JSON 字符串表示的 Deployment 资源
	deploymentJSON := `{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "nginx-example-deployment",
			"namespace": "default"
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "example"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "example"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "nginx",
							"image": "nginx:1.23.4",
							"ports": [
								{
									"containerPort": 80
								}
							]
						}
					]
				}
			}
		}
	}`

	// 解析 JSON 字符串为 Unstructured 对象
	var obj unstructured.Unstructured
	//将Json字符串绑定到obj
	if err := json.Unmarshal([]byte(deploymentJSON), &obj); err != nil {
		fmt.Println("Json解析为Unstructured对象失败： ", err.Error())
		os.Exit(1)
	}
	//fmt.Printf("json:%s\n", &obj.Object)

	// 定义 GroupVersionResource 信息
	gvr := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	// 使用 DynamicClient 创建资源
	resource := dynamicClient.Resource(gvr).Namespace(cf.Namesapce)
	//创建deployment
	_, createErr := resource.Create(context.TODO(), &obj, metav1.CreateOptions{})
	if createErr != nil {
		fmt.Println("Json创建deployment资源失败： ", createErr.Error())
		os.Exit(1)
	}

	fmt.Println("Deployment created successfully")

}
