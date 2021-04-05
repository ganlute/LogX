package common

import (
	"fmt"
	"github.com/ganlute/LogX/data_channel"
	"github.com/ganlute/LogX/data_channel/api"
	"time"
)

const (
	StartTime = "startTime"
)

func GetLogCtx() (ctx map[string]interface{}) {
	ctx = make(map[string]interface{}, 0)
	context, err := getLogContext()
	if err != nil {
		return ctx
	}
	ctx, ok := context.(map[string]interface{})
	if !ok {
		return ctx
	}
	addTimeCost(ctx)
	return ctx
}
func getLogContext() (ctx interface{}, err error) {
	var dataChannel api.DataChannel
	dataChannel = &data_channel.MemMap{}
	return dataChannel.GetContext()
}
func addTimeCost(ctx map[string]interface{}) {
	if startTime, ok := ctx[StartTime]; ok {
		st := startTime.(time.Time)
		timeCost := time.Since(st).Seconds()
		delete(ctx, StartTime)
		ctx["timeCost"] = fmt.Sprintf("%fs", timeCost)
	}
	return
}
