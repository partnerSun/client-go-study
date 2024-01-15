package initcontroller

import cf "goStudy/config"

func init() {
	//metaDatainit()//集群内
	metaDataInit2(cf.MetaNamespace) //集群外

}
