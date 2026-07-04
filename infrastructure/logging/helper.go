package logging

import (
	"runtime"
	"strings"

	"go.uber.org/zap"
)

/*
Contoh runtime func name:
smart-hotel/repository.(*accessRepository).FindOne
*/

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

func NewRepositoryContext() *Context {
	return &Context{layer: LayerRepository}
}

func NewUsecaseContext() *Context {
	return &Context{layer: LayerUsecase}
}

func Start(c *Context, fields ...zap.Field) {
	L().Info("start", append(baseFields(c), fields...)...)
}

func Success(c *Context, id string, fields ...zap.Field) {
	L().Info("success",
		append(baseFields(c), zap.String("id", id))...,
	)
}

func SuccessList(c *Context, count int, fields ...zap.Field) {
	L().Info("success",
		append(baseFields(c), zap.Int("count", count))...,
	)
}

func Failed(c *Context, err error, fields ...zap.Field) {
	L().Error("failed",
		append(baseFields(c), zap.String("error", err.Error()))...,
	)
}

func baseFields(c *Context) []zap.Field {
	return []zap.Field{
		zap.String("layer", string(c.layer)),
		zap.String("caller", callerName()),
	}
}
