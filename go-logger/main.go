package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("This is an info message: user %s, attempt %d", "john_doe", 3)
	// {"level":"info","msg":"This is an info message: user john_doe, attempt 3"}

	// logger := zap.NewExample()
	// logger.Info("This is an info message", zap.String("user", "john_doe"), zap.Int("attempt", 3))
	// {"level":"info","msg":"This is an info message","user":"john_doe","attempt":3}

	// Example
	// logger := zap.NewExample()
	// logger.Info("This message is NewExample logger")

	// Development

	encoder := getEncoderLog()
	writeSyncer := getWriteSyncerLog()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	logger.Info("This is a info message.")
	logger.Error("This is a error message.")
}

// Format log output
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 17161234567 => 2025-06-17T17:16:12.345+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts => time
	encoderConfig.TimeKey = "time"

	// info => INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// Caller
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

// Write log to file
func getWriteSyncerLog() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/log.txt", os.O_RDWR, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stdout)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)

	// return zapcore.AddSync(zapcore.Lock(zapcore.AddSync(&lumberjack.Logger{
	// 	Filename:   "logfile.log",
	// 	MaxSize:    10, // megabytes
	// 	MaxBackups: 3,
	// 	MaxAge:     28,   // days
	// 	Compress:   true, // disabled by default
	// })))
}
