package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"goStudy/utils/logs"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// 常量
const ts = "2006-01-02 15:04:05"

var (
	Namesapce            string
	MetaNamespace        string
	InClusterKubeconfig  *string
	OutClusterKubeconfig *string
	DeploymentName       string
	Replicas             int32
	Image                string
	ClientSet            *kubernetes.Clientset
	DynamicClient        *dynamic.DynamicClient
	Port                 string
	SigningKey           string
	JwtExpireTime        int64 //过期时间 分钟
	Username             string
	Password             string
)

type PodSruct struct {
	Podname   string `json:"podname"`
	Namespace string `json:"namespace"`
	CImage    string `json:"cimage"`
	CName     string `json:"cname"`
}

type NewReturnData struct {
	Status  int
	Message string
}

// 统一日志配置管理
func LogConfig(l string) {
	if l == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: ts}) //json模式
	logrus.SetReportCaller(true)                                    //行号显示
}

func init() {
	//设置命名空间
	viper.SetDefault("name_space", "default") //为空则是所有命名空间
	Namesapce = viper.GetString("name_space")
	//要获取、修改的deployment名字
	viper.SetDefault("deployment_name", "nginx-example-deployment")
	DeploymentName = viper.GetString("deployment_name")

	//设置deployment副本数
	viper.SetDefault("replicas", 3)
	Replicas = viper.GetInt32("replicas")
	//设置image
	viper.SetDefault("image_nameV", "nginx")
	Image = viper.GetString("image_nameV")
	//定义元数据命名空间
	viper.SetDefault("metans", "meta-namespace")
	MetaNamespace = viper.GetString("metans")

	logs.Debug(nil, "init函数启动成功")
	viper.SetDefault("loglevel", "debug")
	viper.SetDefault("port", ":18080")
	//jwt过期时间配置
	viper.SetDefault("Jwt_ExpireTime", 120) //两小时有效期
	//jwt加密secret
	viper.SetDefault("JWT_SigningKey", "this_is_a_siginKey")
	//设置默认用户名密码
	viper.SetDefault("Test_Username", "user1")
	viper.SetDefault("Test_Password", "user1ps")
	viper.AutomaticEnv()
	logl := viper.GetString("loglevel")
	Port = viper.GetString("port")
	JwtExpireTime = viper.GetInt64("Jwt_ExpireTime")
	SigningKey = viper.GetString("JWT_SigningKey")
	//获取用户名密码
	Username = viper.GetString("Test_Username")
	Password = viper.GetString("Test_Password")
	LogConfig(logl)
}
