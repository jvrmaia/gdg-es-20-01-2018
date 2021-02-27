package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Falha ao escrever no log, usando stderr")
	}

	log.WithFields(logrus.Fields{
		"foo": "bar",
		"inc": 10,
	}).Info("Escrevendo no arquivo de log!")
}
