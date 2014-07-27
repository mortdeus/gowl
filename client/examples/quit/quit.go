package main

import (
	gowlclnt "github.com/mortdeus/gowl/client"
)

func main() {
	d := new(gowlclnt.Display)
	if err := d.Connect(""); err != nil {
		panic(err)
	}
	defer d.Disconnect()
	if _, err := d.Write([]byte("quit")); err != nil {
		panic(err)
	}
}
