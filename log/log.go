package log

import "github.com/beego/beego/v2/core/logs"

func init() {
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"app.log","separate":["error", "warning", "info", "debug"]}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(2)
}

func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}

func Debug(f interface{}, v ...interface{}) {
	logs.Debug(f, v...)
}

func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v...)
}

func Warning(f interface{}, v ...interface{}) {
	logs.Warning(f, v...)
}
