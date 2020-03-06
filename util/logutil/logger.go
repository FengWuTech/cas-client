package logutil

import (
	"fmt"
	"time"
)

func prependTime(args []interface{}) []interface{} {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	param := make([]interface{}, 0)
	param = append(param, nowTime)
	param = append(param, args...)
	return param
}

func ErrorF(format string, args ...interface{}) {
	param := prependTime(args)
	fmt.Printf("%v pay-center-client ERROR "+format+"\n", param...)
}

func WarnF(format string, args ...interface{}) {
	param := prependTime(args)
	fmt.Printf("%v pay-center-client WARN "+format+"\n", param...)
}

func InfoF(format string, args ...interface{}) {
	param := prependTime(args)
	fmt.Printf("%v pay-center-client INFO "+format+"\n", param...)
}

func DebugF(format string, args ...interface{}) {
	param := prependTime(args)
	fmt.Printf("%v pay-center-client DEBUG "+format+"\n", param...)
}
