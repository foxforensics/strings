// Carve ASCII and Unicode strings from files.
//
// Usage:
//
//	strings [-nmtao] file
//
// The options are:
//
//	n uint
//	    Minimum string length (default 3).
//	m uint
//	    Maximum string length.
//	t
//		Trim spaces from ends.
//	a
//	    Show ASCII strings only.
//	o
//	    Show file offset.
//
// The arguments are:
//
//	file
//	    File to be carved (required).
package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"go.foxforensics.eu/go-mmap"
	"go.foxforensics.eu/strings/strings"
)

var Usage = `© 2026 Fox Forensics. Licensed under MIT License.
Usage: strings [-nmtao] FILE

  -n  minimum string length
  -m  maximum string length
  -t  trim spaces from ends
  -a  show ASCII strings only
  -o  show file offset

Report bugs at: foxforensics.eu/issues`

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, Usage)
		os.Exit(2)
	}

	x := flag.Uint("n", 3, "minimum string length")
	y := flag.Uint("m", math.MaxUint32, "maximum string length")
	t := flag.Bool("t", false, "trim spaces from ends")
	a := flag.Bool("a", false, "show ASCII strings only")
	o := flag.Bool("o", false, "show file offset")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	f, err := os.Open(flag.Arg(0))

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() { _ = f.Close() }()

	m, err := mmap.Map(f, mmap.RDONLY, 0)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() { _ = m.Unmap() }()

	for s := range strings.Carve(m, *x, *y, *a, *t) {
		if *o {
			_, _ = fmt.Printf("%08x %s\n", s.Offset, s.Value)
		} else {
			_, _ = fmt.Println(s.Value)
		}
	}
}
