package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "goStudy/config"
	"goStudy/middlewares"
	"goStudy/routers"
	"goStudy/utils/jwtutils"
)

func main() {
	//pod.AllPodStatusList(kubeconfig.Namesapce, kubeconfig.ClientSet)
	//pod.PodInfoCheck(kubeconfig.Namesapce, "nginx-55859c47b4-42629",kubeconfig.ClientSet)
	//deployment.CreateDeployment(kubeconfig.ClientSet, cf.DeploymentName)
	//deployment.UpdateDeployment(kubeconfig.ClientSet, cf.DeploymentName)
	//deployment.DelDeployment(kubeconfig.ClientSet, cf.DeploymentName)
	//deployment.CreateDeployByJson(kubeconfig.DynamicClient)
	//cluster.ClusterInfoCheck(kubeconfig.ClientSet)
	//pod.CreatePod(kubeconfig.Namesapce, "nginx-example-pod", kubeconfig.ClientSet)
	//pod.DelPod(kubeconfig.Namesapce, "nginx-example-pod", kubeconfig.ClientSet)
	//pod.GetPod(kubeconfig.Namesapce, "nginx-example-pod", kubeconfig.ClientSet)
	//生成token
	ss, _ := jwtutils.Gentoken("这是一个用户名")
	println("token：", ss)
	claims, err := jwtutils.ParseJwtToken(ss)
	if err != nil {
		fmt.Println("token解析失败:", err.Error())
	} else {
		fmt.Println(claims)                            //打印全部
		fmt.Println(claims.Username)                   // 打印某个字段
		fmt.Println(claims.RegisteredClaims.ExpiresAt) //
	}
	r := gin.New()
	r.Use(middlewares.JwtAuthCheck)
	routers.RegisterRouters(r)
}
