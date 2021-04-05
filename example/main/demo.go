package main

import (
	"github.com/ganlute/LogX/log/common"
	// 1、导包
	"github.com/ganlute/LogX/log/logrus"
	"time"
)

func main() {
	for i := 0; i < 2; i++ {
		// 启动多个协程
		go func(i int) {
			ctx := map[string]interface{}{
				"requestID":      i,
				common.StartTime: time.Now(),
			}
			// 2、每个协程初始化
			logrus.RegisterLogContext(ctx)
			defer logrus.ReleaseLogContext()

			// 3、每个协程直接打印日志
			logrus.Info("hello world")
			Hello1()
		}(i)
	}
	time.Sleep(2 * time.Second)
}
func Hello1() {
	// 3、每个协程直接打印日志
	logrus.Info("hello world 1")
}

// INFO[0000] hello world       caller=".../LogX/example/main/demo.go:23" requestID=0 timeCost=0.000079s
// INFO[0000] hello world       caller=".../LogX/example/main/demo.go:23" requestID=1 timeCost=0.000079s
// INFO[0000] hello world 1     caller=".../LogX/example/main/demo.go:31" requestID=0 timeCost=0.000079s
// INFO[0000] hello world 1     caller=".../LogX/example/main/demo.go:31" requestID=1 timeCost=0.000079s
