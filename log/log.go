package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(app string) {

	Logger = NewLogger("./logs/sofa-server.log", zapcore.InfoLevel, 128, 30, 7, true, true, app)

	Logger.Info("zap logger init success")
}
