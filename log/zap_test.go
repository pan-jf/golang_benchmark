package log

import (
	"testing"

	"go.uber.org/zap"
)

func BenchmarkZap(b *testing.B) {
	logger, _ := zap.NewProduction()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL")
	}
}

func BenchmarkZapSugar(b *testing.B) {
	logger, _ := zap.NewProduction()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar := logger.Sugar()
		sugar.Info("failed to fetch URL")
	}
}
