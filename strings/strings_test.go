package strings

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var bin = filepath.Join("..", "testdata", "test.bin")
var txt = filepath.Join("..", "testdata", "test.txt")

func TestCarve(t *testing.T) {
	for _, tt := range []struct {
		name  string
		path  string
		trim  bool
		ascii bool
		count int
	}{
		{
			name:  "ASCII",
			path:  bin,
			trim:  true,
			ascii: true,
			count: 5590,
		},
		{
			name:  "Unicode",
			path:  txt,
			trim:  false,
			ascii: false,
			count: 614,
		},
	} {
		t.Run("Test Carve "+tt.name, func(t *testing.T) {
			ctx := context.Background()

			buf, err := fixture(tt.path)

			if err != nil {
				t.Fatalf("Carve: %v", err)
			}

			n := 0

			for range Carve(ctx, buf, 3, 255, tt.ascii, tt.trim) {
				n++
			}

			if n != tt.count {
				t.Fatalf("invalid count: %d", n)
			}
		})
	}
}

func BenchmarkCarve(b *testing.B) {
	b.Run("Benchmark Carve", func(b *testing.B) {
		ctx := context.Background()

		bin, err := fixture(bin)

		if err != nil {
			b.Fatalf("Carve: %v", err)
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			for range Carve(ctx, bin, 3, 255, false, false) {
			}
		}
	})
}

func fixture(path string) ([]byte, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	b, err := io.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return b, nil
}
