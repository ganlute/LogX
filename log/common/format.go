package common

import (
	"github.com/ganlute/LogX/data_channel"
	"github.com/ganlute/LogX/data_channel/api"
)

func RegisterLogContext(ctx interface{}) (err error) {
	var dataChannel api.DataChannel
	dataChannel = &data_channel.MemMap{}
	return dataChannel.SetContext(ctx)
}
func ReleaseLogContext() (err error) {
	var dataChannel api.DataChannel
	dataChannel = &data_channel.MemMap{}
	return dataChannel.DeleteContext()
}
func GetLogContext() (ctx interface{}, err error) {
	var dataChannel api.DataChannel
	dataChannel = &data_channel.MemMap{}
	return dataChannel.GetContext()
}