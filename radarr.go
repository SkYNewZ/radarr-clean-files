package main

import (
	"github.com/SkYNewZ/radarr"
	log "github.com/sirupsen/logrus"
)

func deleteMovie(movie *radarr.Movie) error {
	return radarrClient.Movies.Delete(movie, &radarr.DeleteMovieOptions{
		AddExclusion: false,
		DeleteFiles:  true,
	})
}

func getDiskspace() *radarr.Diskspace {
	log.Printf("Get %s free space", moviesPath)
	diskspaces, err := radarrClient.Diskspace.Get()
	if err != nil {
		log.WithFields(log.Fields{
			"go-err": err,
		}).Fatalln("Fail to get diskspace")
	}

	// Filter by /movies
	for _, d := range *diskspaces {
		if d.Path == moviesPath {
			return &d
		}
	}
	log.Fatalf("Fail to get %s free space", moviesPath)
	return nil
}

func initClient() {
	log.WithField("url", radarrURL).Println("Init Radarr client")
	radarrClient, err = radarr.New(radarrURL, radarrAPIKey, nil)
	if err != nil {
		log.WithField("go-err", err).Fatalln("Fail to init Radarr client")
	}
}

func getMovies() radarr.Movies {
	log.Println("Listing movies")
	movies, err := radarrClient.Movies.List()
	if err != nil {
		log.WithField("go-err", err).Fatalln("Fail to list movies")
	}
	log.Printf("Found %d movies", len(movies))
	return movies
}
