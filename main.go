package main

import (
	"goStudy/config"
	cf "goStudy/config"
	"goStudy/deployment"
)

func main() {
	//pod.AllPodStatusList(config.Namesapce, config.ClientSet)
	//pod.PodInfoCheck(config.Namesapce, "nginx-55859c47b4-42629",config.ClientSet)
	//deployment.CreateDeployment(config.ClientSet, cf.DeploymentName)
	deployment.UpdateDeployment(config.ClientSet, cf.DeploymentName)
	//deployment.DelDeployment(config.ClientSet, cf.DeploymentName)
	//deployment.CreateDeployByJson(config.DynamicClient)
	//cluster.ClusterInfoCheck(config.ClientSet)
}
