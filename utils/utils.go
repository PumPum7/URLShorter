package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	file, err := os.OpenFile(Cfg.Logging.Logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to Log to file")
	}
	levels := map[string]logrus.Level{
		"INFO":  logrus.InfoLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
		"FATAL": logrus.FatalLevel,
	}
	Log.SetLevel(levels[Cfg.Logging.LogLevel])
	Log.Info(Log.Level)
}
