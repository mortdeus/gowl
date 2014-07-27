package main

import (
	"bytes"
	"fmt"
	gowlsrv "github.com/mortdeus/gowl/server"
	"net"
	"strings"
)

var done = make(chan int)

func main() {
	d, err := gowlsrv.NewDisplay()
	if err != nil {
		panic(err)
	}
	defer d.Destroy()
	go func() {
		for {
			c, err := d.Accept()
			if err != nil {
				panic(err)
			}
			s, err := read(c)
			if err != nil {
				panic(err)
			}
			if strings.HasPrefix(s, "quit") {
				done <- 0
				return
			}
			fmt.Println(s)
			c.Close()
		}
	}()
	<-done
	fmt.Println("exiting...")
}

func read(c net.Conn) (string, error) {
	b := new(bytes.Buffer)
	if _, err := b.ReadFrom(c); err != nil {
		return "", err
	}
	return b.String(), nil
}
func write(c net.Conn, s string) error {
	if _, err := c.Write([]byte(s)); err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
