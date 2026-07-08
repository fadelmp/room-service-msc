package logging

import (
	"runtime"
	"strings"

	"go.uber.org/zap"
)

func callerName() string {
	for i := 2; i < 10; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			continue
		}

		fn := runtime.FuncForPC(pc).Name()
		if !strings.Contains(fn, "/logging.") {
			parts := strings.Split(fn, "/")
			last := parts[len(parts)-1]

			last = strings.ReplaceAll(last, "(*", "")
			last = strings.ReplaceAll(last, ")", "")

			return last
		}
	}
	return "unknown"
}

/* ===== Public API ===== */

func Success(id string, fields ...zap.Field) {
	L().Info("success", append(baseFields(), zap.String("id", id))...)
}

func SuccessList(count int, fields ...zap.Field) {
	L().Info("success", append(baseFields(), zap.Int("count", count))...)
}

func Failed(err error, fields ...zap.Field) {
	L().Error("failed", append(baseFields(), zap.String("error", err.Error()))...)
}

func baseFields() []zap.Field {
	return []zap.Field{
		zap.String("caller", callerName()),
	}
}
