package gom

// import (
// 	"log"
// 	"os"

// 	logger "github.com/edoger/zkits-logger"
// )

// var _log logger.Logger

// type LogConf struct {
// 	Name     string
// 	Filepath string
// 	Level    string
// }

// func LogInit(conf *LogConf) {
// 	_log = logger.New(conf.Name).
// 		SetDefaultTimeFormat("2006-01-02 15:04:05")

// 	if conf.Filepath != "" {
// 		file, err := os.OpenFile(conf.Filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// defer file.Close()

// 		// 不输出控制台，输出到指定文件
// 		_log.SetOutput(file)
// 	}

// 	if conf.Level == "error" { // 仅打印错误日志，prod
// 		_log.SetLevel(logger.ErrorLevel)
// 	} else if conf.Level == "info" {
// 		_log.SetLevel(logger.InfoLevel)
// 	} else if conf.Level == "debug" {
// 		_log.SetLevel(logger.DebugLevel)
// 	}
// }

// func Log() logger.Logger {
// 	return _log
// }

// func ErrorLog(msg string) {
// 	_log.Error(msg)
// }

// func ErrorfLog(msg string, data ...interface{}) {
// 	_log.Errorf(msg, data...)
// }

// func InfoLog(msg string) {
// 	_log.Info(msg)
// }

// func InfofLog(msg string, data ...interface{}) {
// 	_log.Infof(msg, data...)
// }
// func DebugLog(msg string) {
// 	_log.Debug(msg)
// }

// func DebugfLog(msg string, data ...interface{}) {
// 	_log.Debugf(msg, data...)
// }

// func TraceLog(msg string) {
// 	_log.Trace(msg)
// }

// func TracefLog(msg string, data ...interface{}) {
// 	_log.Tracef(msg, data...)
// }
