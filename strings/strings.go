// Package strings provides methods for ASCII and Unicode string carving.
//
// Source: https://github.com/robpike/strings/blob/master/strings.go
package strings

import (
	"bufio"
	"bytes"
	"context"
	"strconv"
	"strings"
	"unicode/utf8"
)

// String data
type String struct {
	// Offset of string
	Offset uint64
	// Value of string
	Value string
}

// Carve ASCII and Unicode string data.
//
// The returned channel will be closed at the end of the operation.
func Carve(ctx context.Context, data []byte, min, max uint, ascii, trim bool) <-chan *String {
	ch := make(chan *String, 1024)

	go func() {
		defer close(ch)

		b := bufio.NewReader(bytes.NewReader(data))
		s := make([]rune, 0, max)
		i := uint64(0)

		flush := func() bool {
			v := string(s)

			if trim {
				v = strings.TrimSpace(v)
			}

			if uint(utf8.RuneCountInString(v)) >= min {
				select {
				case <-ctx.Done():
					return false
				default:
					ch <- &String{i - uint64(len(v)), v}
				}
			}

			s = s[:0]

			return true
		}

		var r rune
		var n int
		var err error

		for ; ; i += uint64(n) {
			if r, n, err = b.ReadRune(); err != nil {
				flush()
				return
			}

			if !strconv.IsPrint(r) || ascii && (0x20 > r || r > 0x7E) {
				if !flush() {
					return
				}
				continue
			}

			if uint(len(s)) >= max {
				if !flush() {
					return
				}
			}

			s = append(s, r)
		}
	}()

	return ch
}
