package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// API endpoint URL
	url := "https://example.com/api"

	// Open a connection to the database
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
	if err != nil {
		logError(err)
		return
	}
	defer db.Close()

	for {
		// Send a GET request to the API endpoint
		response, err := http.Get(url)
		if err != nil {
			// Log the failure
			logError(err)
			// Wait for 5 minutes before trying again
			time.Sleep(5 * time.Minute)
			continue
		}
		defer response.Body.Close()

		// If the response status code is not 200, log the failure
		if response.StatusCode != http.StatusOK {
			logError(fmt.Errorf("status code: %d", response.StatusCode))
			// Wait for 5 minutes before trying again
			time.Sleep(5 * time.Minute)
			continue
		}

		// Read the response body
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logError(err)
			time.Sleep(5 * time.Minute)
			continue
		}

		// Store the data in the database
		_, err = db.Exec("INSERT INTO data (value) VALUES (?)", data)
		if err != nil {
			logError(err)
			time.Sleep(5 * time.Minute)
			continue
		}

		fmt.Println("Data stored in the database")
		break
	}
}

func logError(err error) {
	// Write the error message to a log file
	f, err := os.OpenFile("api_failures.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(err)
}
