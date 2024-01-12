package initcontroller

import (
	cf "goStudy/config"
	"goStudy/utils"
)

func init() {
	metaDatainit()
	metaDataInit2()

	cf.ClientSet = utils.ClientSetinit(cf.OutClusterKubeconfig)
	cf.DynamicClient = utils.DynamicClientInit(cf.OutClusterKubeconfig)
}
