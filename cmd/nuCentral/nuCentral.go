package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting nuCal")
}

func main() {
	fmt.Println("init")
}
