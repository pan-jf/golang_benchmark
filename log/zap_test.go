package log

import (
	"testing"

	"go.uber.org/zap"
)

// 605.3 ns/op             4 B/op          0 allocs/op
func BenchmarkZap(b *testing.B) {
	logger, _ := zap.NewProduction()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL")
	}
}

//  615.4 ns/op             4 B/op          0 allocs/op
func BenchmarkZapSugar(b *testing.B) {
	logger, _ := zap.NewProduction()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sugar := logger.Sugar()
		sugar.Info("failed to fetch URL")
	}
}
