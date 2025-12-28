package fetchdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func WebFetcher(fetchData chan string, url string, wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in WebFetcher for %s: %v\n", url, r)
		}
	}()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		fetchData <- fmt.Sprintf("ERROR: %s - %v", url, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading body for %s: %v\n", url, err)
		fetchData <- fmt.Sprintf("ERROR reading body: %s", url)
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		fetchData <- fmt.Sprintf("RAW DATA from %s: %s", url, string(body))
		return
	}

	fetchData <- fmt.Sprintf("FETCHED from %s:\n%s", url, prettyJSON.String())
}
