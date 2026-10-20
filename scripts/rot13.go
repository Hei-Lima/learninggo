package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(s byte) byte {
	c := s
	if s >= 'a' && s <= 'z' {
		c = byte('a' + ((s-'a')+13)%26)
	} else if s >= 'A' && s <= 'Z' {
		c = byte('A' + ((s-'A')+13)%26)
	}
	return byte(c)
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	s := 0
	n, e := rot.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
		s++
	}

	return s, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
