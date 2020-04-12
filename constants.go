package main

import (
	"net/http"
	"os"

	"github.com/SkYNewZ/radarr"
	log "github.com/sirupsen/logrus"
)

var (
	radarrClient      *radarr.Service
	httpClient        *http.Client = http.DefaultClient
	radarrURL         string
	radarrAPIKey      string
	healthChecksIOURL string
	deletedMovies     int = 0
	errs              int = 0
	err               error
)

const (
	wantedFreeSpace      int64  = 100000000000
	moviesPath           string = "/movies"
	maximumDeletedMovies int    = 10
)

func initConfig() {
	// Read Radarr conf from env
	var neededVars [3]string = [3]string{"RADARR_URL", "RADARR_API_KEY", "HEALTHCHECKS_IO_URL"}

	for i, v := range neededVars {
		if ok := os.Getenv(v); ok == "" {
			log.Fatalf("%s is missing", v)
		}
		switch i {
		case 0:
			radarrURL = os.Getenv(v)
		case 1:
			radarrAPIKey = os.Getenv(v)
		case 2:
			healthChecksIOURL = os.Getenv(v)
		}
	}
}
