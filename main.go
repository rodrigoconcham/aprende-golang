package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	duracion := time.Since(start)
	fmt.Printf("Fetched %s en %v - Status %s\n", url, duracion, resp.Status)
}

func main() {

	urls := []string{
		"https://rickandmortyapi.com/api/character/1",
		"https://rickandmortyapi.com/api/character/2",
		"https://rickandmortyapi.com/api/character/3",
		"https://rickandmortyapi.com/api/character/4",
	}

	start := time.Now()

	for _, url := range urls {
		fetchURL(url)
	}

	duracionTotal := time.Since(start)
	fmt.Printf("Tiempo total de la secuencia: %v\n", duracionTotal)

}
