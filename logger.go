package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type logger struct{}

func (l *logger) Format(entry *log.Entry) ([]byte, error) {
	m := fmt.Sprintf("[%s] [%s] %s", entry.Time.Format(time.RFC822), entry.Level, entry.Message)
	for k, v := range entry.Data {
		m += fmt.Sprintf(` %s="%v"`, k, v)
	}
	m += "\n"
	return []byte(m), nil
}

func init() {
	if l := os.Getenv("LOG_LEVEL"); l != "" {
		level, err := log.ParseLevel(l)
		if err != nil {
			log.Fatalln(err)
		}
		log.SetLevel(level)
	}
	log.SetFormatter(new(logger))
}
