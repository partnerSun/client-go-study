package pod

import (
	cf "client-go-study/config"
	_ "client-go-study/controllers/initcontroller"
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func GetPod(ns string, podName string) (*corev1.Pod, error) {
	// 使用 Pod 名称和命名空间查询指定 Pod 的详细信息
	pod, err := cf.ClientSet.CoreV1().Pods(ns).Get(context.TODO(), podName, metav1.GetOptions{})
	//podStatus := pod.Status.ContainerStatuses[0].State.Waiting.Reason
	if err != nil {
		//fmt.Printf("Error! failed getting Pod %s in namespace %s: %s\n", podName, ns, err.Error())
		return nil, err
		//os.Exit(1)
	}
	fmt.Printf("Getting Pod  %s in namespace %s\n", podName, ns)
	return pod, nil
}

func CreatePod(cpinfo cf.PodSruct) (*corev1.Pod, error) {
	podName := cpinfo.Podname
	ns := cpinfo.Namespace
	//cName := cpinfo["cname"]
	cImage := cpinfo.CImage
	cName := cpinfo.CName
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ns,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  cName,
					Image: cImage,
				},
			},
		},
	}
	_, err := cf.ClientSet.CoreV1().Pods(ns).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error! failed to Create Pod %s in namespace %s: %s", podName, ns, err.Error())
		//os.Exit(1)
		return nil, err
	}
	fmt.Printf("Create Pod  %s in namespace %s success!", podName, ns)
	return nil, nil
}

func PodInfoCheck(ns string, podName string) {

	// 使用 Pod 名称和命名空间查询指定 Pod 的详细信息
	pod, _ := GetPod(ns, podName)
	// 打印获取到的 Pod 详细信息
	fmt.Printf("Pod Details:\n")
	fmt.Printf("Name: %s\n", pod.Name)
	fmt.Printf("Namespace: %s\n", pod.Namespace)
	fmt.Printf("State: %s\n", pod.Status.Phase)
	//fmt.Printf("Conditions: %s\n", pod.Status.Conditions)
	//fmt.Printf("Message: %s\n", pod.Status.Message)
	//fmt.Printf("Reason: %s\n", pod.Status.Reason)
	//fmt.Printf("NominatedNodeName: %s\n", pod.Status.NominatedNodeName)
	fmt.Printf("HostIP: %s\n", pod.Status.HostIP)
	fmt.Printf("HostIPs: %s\n", pod.Status.HostIPs)
	fmt.Printf("PodIP: %s\n", pod.Status.PodIP)
	fmt.Printf("PodIPs: %s\n", pod.Status.PodIPs)
	fmt.Printf("InitContainerStatuses: %s\n", pod.Status.InitContainerStatuses)
	//fmt.Printf("ContainerStatuses: %s\n", pod.Status.ContainerStatuses)
	//fmt.Printf("QOSClass: %s\n", pod.Status.QOSClass)
	fmt.Printf("NodeName: %s\n", pod.Spec.NodeName)
	//fmt.Printf("Volumes: %s\n", pod.Spec.Volumes)
	fmt.Printf("Affinity: %s\n", pod.Spec.Affinity)
	fmt.Printf("DNSPolicy: %s\n", pod.Spec.DNSPolicy)
	fmt.Printf("HostNetwork: %s\n", pod.Spec.HostNetwork)
	//fmt.Printf("Hostname: %s\n", pod.Spec.Hostname)
	fmt.Printf("Containers: %s\n", pod.Spec.Containers)
	fmt.Printf("RestartPolicy: %s\n", pod.Spec.RestartPolicy)
	fmt.Printf("NodeSelector: %s\n", pod.Spec.NodeSelector)

	// 还可以打印更多 Pod 的其他信息，如标签、容器信息、IP 等
	// 你可以根据实际需求，进一步处理获取到的 Pod 详细信息

}
func DelPod(ns string, podName string) (*corev1.Pod, error) {
	_, err := GetPod(ns, podName)
	if err != nil {
		fmt.Printf("Error! failed getting Pod %s in namespace %s: %s\n", podName, ns, err.Error())
		return nil, err
		os.Exit(1)
		//os.Exit(1)
	}
	fmt.Printf("Getting Pod  %s in namespace %s", podName, ns)

	err = cf.ClientSet.CoreV1().Pods(ns).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Printf("Error! failed to Delele pod %s:%s", podName, err.Error())
		return nil, err
		os.Exit(1)
	}
	fmt.Printf("Delele pod %s success", podName)
	return nil, nil
}

func AllPodStatusList(ns string) (*corev1.PodList, error) {

	// 查询 Pod 列表
	pods, err := cf.ClientSet.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error getting Pod list: %v\n", err)
		os.Exit(1)
		return nil, err
	}
	// 打印 Pod 名称
	fmt.Printf("Pods in namespace %s:\n", ns)
	for _, pod := range pods.Items {
		fmt.Printf(" - Name: %s, State: %s\n", pod.Name, pod.Status.Phase)
		//return pod, err
	}
	return pods, nil
}
