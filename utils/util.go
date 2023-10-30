package utils

import (
	"bebecare-go-api-1/models/invite"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func GetDate() string {
	currentTime := time.Now()

	return currentTime.Format("2006-01-02 15:04:05")

}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func PrintConsoleLog(funcName interface{}, a ...interface{}) {
	fmt.Println("date :", GetDate(), "::: func :", GetFunctionName(funcName), a)
}

func GenerateRandomString(length int) string {
	var out strings.Builder
	charSet := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		out.WriteString(string(randomChar))
	}
	return out.String()
}

func Base64Encoding(in string) string {
	if in == "" {
		return in
	}
	out := base64.StdEncoding.EncodeToString([]byte(in))

	return out
}

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func IsDev() bool {
	hostname, _ := os.Hostname()
	if strings.Contains(hostname, "bebecare-real") {
		return false
	}

	return true
}

func MakeInviteCodeUnique() int {
	var rangeMin = 100000
	var rangeMax = 999999
	var inviteCode = rangeMin + rand.Intn(rangeMax-rangeMin)

	// true = 중복, false = 중복안됨
	checkDupCode, err := inviteModel.DupInviteCode(inviteCode)
	if err != nil {
		return 0
	}

	// 중복이면 재귀
	if checkDupCode == true {
		return MakeInviteCodeUnique()
	}

	return inviteCode
}
