package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// 50959 ns/op             641 B/op         17 allocs/op
func BenchmarkLogrus(b *testing.B) {
	var log = logrus.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("failed to fetch URL")
	}
}
