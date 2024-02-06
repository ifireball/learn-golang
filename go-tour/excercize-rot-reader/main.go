package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buf []byte) (int, error) {
	inp := make([]byte, len(buf))
	nread, err := r.r.Read(inp)
	if err != nil {
		return 0, err
	}
	for i, char := range inp {
		if char >= 'a' && char <= 'z' {
			buf[i] = (char-'a'+13)%26 + 'a'
		} else if char >= 'A' && char <= 'Z' {
			buf[i] = (char-'A'+13)%26 + 'A'
		} else {
			buf[i] = char
		}
	}
	return nread, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
