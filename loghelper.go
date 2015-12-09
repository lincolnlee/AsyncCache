package AsyncCache

import (
	"github.com/astaxie/beego/logs"
)

type loghelper struct {
	Logger *logs.BeeLogger
}

func newLogger() *logs.BeeLogger {
	log := logs.NewLogger(10000)
	log.SetLogger("console", `{"level":6}`)
	return log
}
