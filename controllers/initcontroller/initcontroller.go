package initcontroller

import cf "goStudy/config"

// clientset创建唯一入口
func init() {
	//metaDatainit()//集群内
	metaDataInit2(cf.MetaNamespace) //集群外

}
