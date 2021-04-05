package logrus

import (
	"github.com/ganlute/LogX/log/common"
)

func RegisterLogContext(ctx interface{}) (err error) {
	return common.RegisterLogContext(ctx)
}
func ReleaseLogContext() (err error) {
	return common.ReleaseLogContext()
}
func GetLogContext() (ctx interface{}, err error) {
	return common.GetLogContext()
}
