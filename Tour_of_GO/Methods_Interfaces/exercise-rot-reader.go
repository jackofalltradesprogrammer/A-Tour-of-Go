package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	n, e := rot13.r.Read(b)
	for x := range b {
		b[x] = r13(b[x])
	}
	return n, e

}

func r13(b byte) byte {
	var c byte
	switch {
	case b >= 'A' && b <= 'Z':
		c = (b-'A'+13)%26 + 'A'
	case b >= 'a' && b <= 'z':
		c = (b-'a'+13)%26 + 'a'
	default:
		c = b
	}
	return c
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
