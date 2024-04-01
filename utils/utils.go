package utils

import (
	"client-go-study/utils/logs"
	"k8s.io/apimachinery/pkg/util/json"
)

func Struct2map(s interface{}) map[string]string {
	j, err := json.Marshal(s) //将结构体转换为json
	if err != nil {
		logs.Error(nil, "结构体转换为map失败")
	}
	m := make(map[string]string)
	json.Unmarshal(j, &m)

	return m
}
