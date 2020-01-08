package base

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Log *zap.SugaredLogger

const (
	output_dir = "./logs/"
	out_path   = output_dir + "foot.log"
	err_path   = output_dir + "foot.err"
)

func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func init() {
	_, err := os.Stat(output_dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(output_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
		}
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "console" /* json console */,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "t",
			LevelKey:       "level",
			NameKey:        "log",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout", out_path},
		ErrorOutputPaths: []string{"stderr", err_path},
		//InitialFields: map[string]interface{}{
		//	"app": "foot",
		//},
	}

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	Log = logger.Sugar()
	defer logger.Sync()
}
