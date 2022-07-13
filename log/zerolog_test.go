package log

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func BenchmarkZerolog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Info().Msg("failed to fetch URL")
	}
}
