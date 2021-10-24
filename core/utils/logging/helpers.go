package logging

import (
	"io/ioutil"
	"log"
	"os"

	"gitlab.com/Dank-del/SibylAPI-Go/core/utils/timeUtils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SUGARED *zap.SugaredLogger

func InitZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, _ := config.Build(zap.AddCallerSkip(1))
	return logger
}

func LoadLogger() func() {
	if SUGARED != nil {
		return func() {}
	}
	loggerMgr := InitZapLog()
	zap.ReplaceGlobals(loggerMgr)
	SUGARED = loggerMgr.Sugar()

	return func() {
		err := loggerMgr.Sync()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func Warn(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Warn(args...)
	} else {
		log.Println(args...)
	}
}

func Error(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Error(args...)
	} else {
		log.Println(args...)
	}
}

func Info(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Info(args...)
	} else {
		log.Println(args...)
	}
}

func Infof(template string, args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Infof(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Debug(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debug(args...)
	} else {
		log.Println(args...)
	}
}

func Debugf(template string, args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Debugf(template, args...)
	} else {
		log.Printf(template, args...)
	}
}

func Fatal(args ...interface{}) {
	if SUGARED != nil {
		SUGARED.Fatal(args...)
	} else {
		log.Fatal(args...)
	}
}

func LogPanic(details []byte) {
	p := string(os.PathSeparator)
	_ = ioutil.WriteFile("logs"+p+"panics"+
		"panic_"+timeUtils.GenerateCurrentDateTime(),
		details, 0)
}
