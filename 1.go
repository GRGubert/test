package main

import (
	"fmt"
	"log"

	sl "github.com/vpatsenko/skilllog"
)

func main() {
	l, err := sl.NewLogger("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(l)
	l.Info("some info log")
	l.Error(fmt.Errorf("some error"))
	l.Warn("some warn")
	if err := l.Close(); err != nil {
		log.Fatalln(err)
	}
}
