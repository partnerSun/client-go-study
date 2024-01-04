package deployment

import (
	"context"
	"fmt"
	cf "goStudy/config"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
	"os"
)

func int32Ptr(i int32) *int32 {
	return &i
}

func CreateDeployment(client *kubernetes.Clientset, deploymentName string) {
	var deployment *v1.Deployment
	deployment = &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: cf.Namesapce,
			Labels: map[string]string{
				"app": cf.DeploymentName,
			},
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(cf.Replicas),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": cf.DeploymentName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": cf.DeploymentName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  cf.DeploymentName,
						Image: cf.Image,
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
						}},
					}},
				},
			},
		},
	}

	_, err := client.AppsV1().Deployments(cf.Namesapce).Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		//return fmt.Errorf("failed to create deployment: %w", err)
		fmt.Printf("Error! failed to create deployment: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Created deployment %q in namespace %q\n", cf.DeploymentName, cf.Namesapce)
}

func GetDeployment(client *kubernetes.Clientset, deploymentName string) (*v1.Deployment, error) {
	//deploymentsClient := cluster.AppsV1().Deployments(cf.Namesapce)

	// 获取当前的 Deployment 对象
	deployment, err := client.AppsV1().Deployments(cf.Namesapce).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		//return fmt.Errorf("failed to create deployment: %w", err)
		return nil, err

	}
	fmt.Printf("GET deployment %q success\n", cf.DeploymentName)
	return deployment, nil
}

func UpdateDeployment(client *kubernetes.Clientset, deploymentName string) {
	//deploymentsClient := cluster.AppsV1().Deployments(cf.Namesapce)

	// 通过GetDeployment判断是否存在deployment，并且返回deployment实例对象
	deployment, err := GetDeployment(client, deploymentName)
	if err != nil {
		//return fmt.Errorf("failed to create deployment: %w", err)
		fmt.Printf("Error! failed to get deployment: %s", err.Error())
		os.Exit(1)
	}

	// 修改 Deployment 对象的一些属性,
	// 注意！！：有些属性如果不存在，直接修改会报空指针
	// 修改 副本数
	deployment.Spec.Replicas = int32Ptr(4) // 更新副本数为 4
	//修改命名空间,命名空间不存在会报错
	deployment.ObjectMeta.Namespace = "default"
	//修改镜像版本
	deployment.Spec.Template.Spec.Containers[0].Image = "nginx:1.24.0"
	//修改容器启动端口
	deployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort = 8080
	//修改dns解析策略
	deployment.Spec.Template.Spec.DNSPolicy = "ClusterFirst"
	//修改重启策略
	deployment.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyAlways
	//修改label,修改需要添加老标签，只会保留当前传的键值
	deployment.Labels["a"] = "b"

	// 修改资源限制（假设需要设置 CPU 和内存限制）
	deployment.Spec.Template.Spec.Containers[0].Resources = corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("1"),
			corev1.ResourceMemory: resource.MustParse("1Gi"),
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse("500m"),
			corev1.ResourceMemory: resource.MustParse("200Mi"),
		},
	}
	// 修改 Rolling Update 策略（假设需要修改 maxSurge 和 maxUnavailable）
	deployment.Spec.Strategy.RollingUpdate = &v1.RollingUpdateDeployment{
		MaxSurge: &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},

		MaxUnavailable: &intstr.IntOrString{Type: intstr.String, StrVal: "25%"},
		//或者,使用整数
		//MaxUnavailable: &intstr.IntOrString{Type: intstr.Int, IntVal: 25},
	}

	// 更新 Deployment 对象
	// 使用 retry 来处理更新的重试机制
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, updateErr := client.AppsV1().Deployments(cf.Namesapce).Update(context.TODO(), deployment, metav1.UpdateOptions{})
		return updateErr
	})
	if err != nil {
		//return fmt.Errorf("failed to create deployment: %w", err)
		fmt.Printf("failed to create deployment: ", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Updated deployment %q in namespace %q success!", cf.DeploymentName, cf.Namesapce)
}

func DelDeployment(client *kubernetes.Clientset, deploymentName string) {
	//deploymentsClient := cluster.AppsV1().Deployments(cf.Namesapce)

	// 获取当前的 Deployment 对象
	_, err := GetDeployment(client, deploymentName)
	if err != nil {
		//return fmt.Errorf("failed to create deployment: %w", err)
		fmt.Printf("Error! failed to get deployment: %s", err.Error())
		os.Exit(1)
	}

	// 使用 retry 来处理更新的重试机制
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		updateErr := client.AppsV1().Deployments(cf.Namesapce).Delete(context.TODO(), cf.DeploymentName, metav1.DeleteOptions{})
		return updateErr
	})
	if err != nil {
		//return fmt.Errorf("failed to create deployment: %w", err)
		fmt.Printf("Error! failed to delete deployment: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Delete deployment %q in namespace %q success!", cf.DeploymentName, cf.Namesapce)
}
