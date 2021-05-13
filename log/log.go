package log

import (
	"github.com/sirupsen/logrus"
)

var (
	defaultLogger *logrus.Logger = nil
)

func Logger() *logrus.Logger {

	if defaultLogger == nil {
		defaultLogger = logrus.New()
		defaultLogger.SetReportCaller(true)
		defaultLogger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true, //DisableColors: true,
		})
		defaultLogger.AddHook(&ModuleNameHook{})

	}

	return defaultLogger
}
