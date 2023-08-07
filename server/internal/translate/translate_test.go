package translate

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	tr := Tr("zh_CN", "backend")
	t.Log(tr)
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		
	}
}
