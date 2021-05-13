package log

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

//Hook de adiciona  o nome do modulo antes do log
type ModuleNameHook struct {
}

func (h *ModuleNameHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *ModuleNameHook) Fire(e *logrus.Entry) error {

	appName := "undefined"

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		appName = buildInfo.Main.Path[strings.LastIndex(buildInfo.Main.Path, "/")+1:] //nome do modulo
		appName += " " + buildInfo.Main.Version                                       //versao do modulo

	} else {
		fmt.Println("Nao pode ler a build info")
	}

	e.Message = "[" + appName + "] " + e.Message
	//e.Data["app"] = appName"
	return nil
}
