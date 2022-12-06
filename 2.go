package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		l.Sync()
	}()
	l.Info("some info")
}
