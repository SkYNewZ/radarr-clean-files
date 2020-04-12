package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func start() {
	u := fmt.Sprintf("%s/start", healthChecksIOURL)
	log.WithField("healthchecks_io_url", u).Println("Signal start")
	_, err := httpClient.Head(u)
	if err != nil {
		log.WithField("go-err", err).Fatalln("Fail to ping start URL")
	}
}

func stop() {
	var u string = healthChecksIOURL
	switch errs > 0 {
	case true:
		u += "/fail"
		log.WithField("healthchecks_io_url", u).Println("Signal failure")
	default:
		log.WithField("healthchecks_io_url", u).Println("Signal stop")
	}

	_, err := httpClient.Head(u)
	if err != nil {
		log.WithField("go-err", err).Fatalln("Fail to ping start URL")
	}
}
