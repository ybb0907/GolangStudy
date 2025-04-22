// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// type IPAddr [4]byte

// // TODO: 为 IPAddr 添加一个 "String() string" 方法。
// func (ip *IPAddr) String() string {
// 	var s string = ip.String()
// 	return s
// }

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for name, ip := range hosts {
// 		fmt.Printf("%v: %v\n", name, ip)
// 	}
// }

// func Sqrt(x float64) (float64, error) {
// 	return 0, nil
// }

// type MyReader struct{}

// func (m *MyReader) Read(p []byte) (int, error) {
// 	for i := range p {
// 		p[i] = 'p'
// 	}
// 	return len(p), nil
// }

// type rot13Reader struct {
// 	r io.Reader
// }

// // func (m rot13Reader) Read(p []byte) (int, error) {

// // }

// func test1() {
// 	s := strings.NewReader("Lbh penpxrq gur pbqr!")
// 	r := rot13Reader{s}
// 	io.Copy(os.Stdout, &r)
// }
