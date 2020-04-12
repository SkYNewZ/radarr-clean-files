/*
Simple script to delete older files from Radarr. Can be deleted on disk too
*/
package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	now := time.Now()
	initConfig()

	// healthchecks.io
	start()
	defer stop()

	// Create Radarr client
	initClient()

	// Get disk space
	diskspace := getDiskspace()

	var x float64 = float64(diskspace.FreeSpace) / 1e12
	log.Printf("%.2f TiB left on disk", x)

	// Test if diskspace is < 100GB
	if diskspace.FreeSpace > wantedFreeSpace {
		log.Println("Nothing to do, diskpace is sufficient.")
		return
	}

	// Get movies
	movies := getMovies()

	// Search and process movies
	deadLineToKeepFiles := time.Now().Add(-730 * time.Hour) // 1 month
	log.Printf("Searching movies older than %s", deadLineToKeepFiles.Format(time.RFC822))

	// Main process
	for _, movie := range movies {
		// Limit to 10 movies
		if deletedMovies >= maximumDeletedMovies {
			log.Println("Maximum number of deleted films reached")
			break
		}

		if !movie.Monitored {
			// We search only monitored movies
			log.WithFields(log.Fields{
				"title": movie.Title,
				"id":    movie.ID,
				"added": movie.Added.Format(time.RFC3339),
			}).Debugln("Not monitored")
			continue
		}

		if !movie.Downloaded {
			// We search only downloaded movies
			log.WithFields(log.Fields{
				"title": movie.Title,
				"id":    movie.ID,
				"added": movie.Added.Format(time.RFC3339),
			}).Debugln("Not downloaded")
			continue
		}

		if !movie.MovieFile.DateAdded.Before(deadLineToKeepFiles) {
			log.WithFields(log.Fields{
				"title": movie.Title,
				"id":    movie.ID,
				"added": movie.Added.Format(time.RFC3339),
			}).Debugln("Too recent")
			continue
		}

		// Delete it
		log.WithFields(log.Fields{
			"title": movie.Title,
			"id":    movie.ID,
		}).Println("Deleting movie")

		if err := deleteMovie(movie); err != nil {
			log.WithFields(log.Fields{
				"go-err": err,
				"title":  movie.Title,
				"id":     movie.ID,
			}).Errorln("Fail to delete movie")
			errs++
			continue
		}
		deletedMovies++
	}

	elapsed := time.Since(now)
	log.Printf("Delete %d movies", deletedMovies)
	log.Printf("Execution took %s", elapsed.Round(time.Second))
}
