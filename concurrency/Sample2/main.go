package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	mut.Lock()
	defer mut.Unlock()
	fmt.Printf("[%d] %s\n", res.StatusCode, url)

}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main  <url1> <url2> ... <urln> ")
	}

	for _, url := range os.Args[1:] {
		go sendRequest("https://" + url)
		wg.Add(1)
	}

	wg.Wait()

}

// benchmarks
// 1. 0.77s user 0.37s system 19% cpu 5.804 total
// 2. 0.78s user 0.40s system 14% cpu 7.986 total
// 3. 0.82s user 0.39s system 46% cpu 2.602 total
// 4. 0.78s user 0.38s system 84% cpu 1.380 total
// 5. 0.78s user 0.39s system 68% cpu 1.692 total
// 6. 0.80s user 0.37s system 81% cpu 1.428 total
