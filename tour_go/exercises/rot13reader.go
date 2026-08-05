package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rtr rot13Reader) Read(b []byte) (int, error) {
	n, res := rtr.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return n, res
}

func rot13(b byte) byte {
    switch {
    case b >= 'A' && b <= 'Z':
        return 'A' + (b - 'A' + 13) % 26
    case b >= 'a' && b <= 'z':
        return 'a' + (b - 'a' + 13) % 26
    default:
        return b
    }
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
