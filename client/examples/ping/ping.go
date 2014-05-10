package main

import (
	gowl "github.com/mortdeus/gowl/client"
	"log"
	"time"
)

func main() {
	if err := gowl.Display.Connect(""); err != nil {
		log.Fatal(err)
	}
	c := make(chan int)
	go func(c chan int) {
		if err := gowl.Display.Read(); err != nil {
			log.Fatal(err)
		}
		c <- 1
	}(c)
	f, err := gowl.Display.File()
	if err != nil {
		log.Fatal(err)
	}
	select {
	case <-time.Tick(time.Second * 10):
		log.Fatal("timeout\n")
	case <-c:
		log.Print(f.Name())
		log.Print("success")
	}
}
