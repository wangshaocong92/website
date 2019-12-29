package util

import (
	"github.com/jander/golog/logger"
	"log"
	"time"
)

const (
	DEBUG  = iota
	INFO
	WARN
	ERROR
)

func init() {
	updateHander()
	// logger set flags
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// logger set log level
	logger.SetLevel(logger.INFO)
	go timer()
}

func updateHander() {
	Close()
	currenttiem := time.Now().Format("2006-01-02")
	hander := logger.NewFileHandler(log_dir()+currenttiem+".log")
	logger.SetHandlers(hander)
}

func Close()  {
	logger.Close()
}

func Log(level int,message ...interface{}){
	switch level {
	case DEBUG:
		logger.Debug(message,"\r\n")
		break
	case INFO:
		logger.Info(message,"\r\n")
		break
	case WARN:
		logger.Warn(message,"\r\n")
		break
	case ERROR:
		logger.Error(message,"\r\n")
		break
	default:
		Log(ERROR,"日志级别错误，请确认代码")
	}
}

func timer() {
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
        updateHander()
	}
}

func log_dir()string{
	if LogDir==""{
		return "../log/"
	}else{
		if string(LogDir[len(LogDir)-1:])=="/"{
			return LogDir
		}else{
			return LogDir + "/"
		}
	}
}