package main

import (
	"github.com/mediocregopher/radix.v2/pool"
	"log"
	"redsync"
)

func main() {
	p, err :=pool.New("tcp", "localhost:6379", 1)
	if(err != nil) {
		log.Fatal(err)
	}
	red := redsync.New([] redsync.Pool{redsync.RedixV2Pool{p}})
	m := red.NewMutex("foo1")
	log.Println(m.Lock())
	log.Println(m.Unlock())


}
