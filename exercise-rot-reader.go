package main
// cheated from: https://qiita.com/rock619/items/f412d1f870a022c142d0
import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
// at first I did not know the concept of rot13
func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err:= rot13.r.Read(b)
}
// the above is the same as follows:
// func (rot13 rot13Reader) Read(b []byte) (n int, err error) {
// 	n, err = rot13.r.Read(b)
// }
	for i := range b {
		if b[i] >= 'A' && b[i] <= 'M' || b[i] >= 'a' && b[i] <= 'm' {
			b[i] += 13
		} else if b[i] >= 'N' && b[i] <= 'Z' || b[i] >= 'n' && b[i] <= 'z' {
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
