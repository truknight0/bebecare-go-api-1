package log

import (
	"bebecare-go-api-1/beans/constants"
	"bebecare-go-api-1/utils/config"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var logLevel constants.LOG_LEVEL

func init() {
	serviceHome := os.Getenv("bebecare go api 1")
	if serviceHome == "" {
		serviceHome = ".."
	}
	logPath := fmt.Sprintf("%s/logs/%s", serviceHome, "bebecare-go-api-1_%Y%m%d.log")
	rl, _ := rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(time.Hour*24*30),
	)
	log.SetOutput(rl)

	logLevel = constants.LOG_LEVEL(config.GetIntDefault("log.level", 2))
}

func ERROR(msg string) {
	if logLevel < constants.LOG_LEVEL_ERROR {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	file = parseFilePath(file)
	log.Println(fmt.Sprintf("[ERROR] %s:%d %s", file, line, msg))
}

func INFO(msg string) {
	if logLevel < constants.LOG_LEVEL_INFO {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	file = parseFilePath(file)
	log.Println(fmt.Sprintf("[INFO] %s:%d %s", file, line, msg))
}

func DEBUG(msg string) {
	if logLevel < constants.LOG_LEVEL_DEBUG {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	file = parseFilePath(file)
	log.Println(fmt.Sprintf("[DEBUG] %s:%d %s", file, line, msg))
}

func parseFilePath(fPath string) (fileName string) {
	_, fileName = filepath.Split(fPath)

	return
}
