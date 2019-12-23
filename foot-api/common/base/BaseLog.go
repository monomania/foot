package base

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Log *logrus.Logger

const LogFile_Path = "./foot.log"

func init() {
	logFile, err := os.OpenFile(LogFile_Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer logFile.Close()
	writers := []io.Writer{
		logFile,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	Log = logrus.New()
	Log.SetOutput(fileAndStdoutWriter)
	Log.Info("foot初始化...")
}
