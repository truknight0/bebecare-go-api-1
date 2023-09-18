package config

import (
	"fmt"
	"gopkg.in/go-ini/ini.v1"
	"os"
	"strings"
)

var section *ini.Section
var serviceMode string

func init() {
	//환경 변수
	f, err := ini.Load("utils/config/bebecare-go-api-1.ini")
	if err != nil {
		panic(err)
	}
	//section 가져 오기(hostname 에 'bebecare-real' 이 포함되어 있으면 product)
	serviceMode = "dev"
	hostname, _ := os.Hostname()
	if strings.Contains(hostname, "bebecare-real") {
		serviceMode = "product"
	}
	fmt.Println("=============================================")
	fmt.Println("  bebecare go api 1 configuration  ")
	fmt.Println("=============================================")
	section = f.Section(serviceMode)
	keys := section.Keys()
	for _, key := range keys {
		fmt.Println(key.Name() + "=" + key.Value())
	}
}

func IsDevMode() bool {
	return serviceMode == "dev"
}

func GetString(k string) string {
	return GetStringDefault(k, "")
}

func GetStringDefault(k, d string) string {
	return section.Key(k).MustString(d)
}

func GetInt(k string) int {
	return GetIntDefault(k, 0)
}

func GetIntDefault(k string, d int) int {
	return section.Key(k).MustInt(d)
}

func GetBool(k string) bool {
	return GetBoolDefault(k, false)
}

func GetBoolDefault(k string, d bool) bool {
	return section.Key(k).MustBool(d)
}
