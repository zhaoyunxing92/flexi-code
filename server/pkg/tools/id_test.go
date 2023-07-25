package tools

import (
	"testing"
)

func BenchmarkCreateSpaceId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreateSpaceId()
		//b.Log(spaceId)
	}
}
