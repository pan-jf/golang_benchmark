package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func BenchmarkLogrus(b *testing.B) {
	var log = logrus.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info("failed to fetch URL")
	}
}
