package logrus

import (
	"github.com/ganlute/LogX/log/common"
	"github.com/ganlute/LogX/utils"
	log "github.com/sirupsen/logrus"
)

func getLogEntry(n int) *log.Entry {
	n++
	fields := log.Fields{}
	fields["caller"] = utils.GetCallerName(n)
	ctx := common.GetLogCtx()
	for key, value := range ctx {
		fields[key] = value
	}
	return log.WithFields(fields)
}
