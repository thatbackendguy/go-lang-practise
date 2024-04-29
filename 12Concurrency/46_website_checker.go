package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func checkAndSave(url string, wg *sync.WaitGroup) {

	resp, err := http.Get(url)

	if err != nil {

		log.Fatal(err)

	} else {

		if resp.StatusCode == http.StatusOK {

			bodyBytes, err := io.ReadAll(resp.Body)

			if err != nil {

				log.Fatal(err)

			} else {
				fileName := strings.Split(url, "//")[1]

				fileName += ".html"

				file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)

				defer file.Close()

				if err != nil {

					log.Fatal(err)

				} else {

					bufferedWriter := bufio.NewWriter(file)

					_, err := bufferedWriter.Write(bodyBytes)

					// error handling

					if err != nil {

						log.Fatal(err)
					}

					bufferedWriter.Flush()
				}

				//err := ioutil.WriteFile(url+".txt", bodyBytes, 0664)

			}

		} else if resp.StatusCode == http.StatusNotFound {

			log.Printf("%s not found", url)

		} else {

			log.Printf("%s is down", url)

		}
	}

	wg.Done()

}

func main() {

	var wg sync.WaitGroup

	urls := []string{"https://www.google.com", "https://golang.org", "https://www.thatbackendguy.com"}

	wg.Add(len(urls))

	for _, url := range urls {

		go checkAndSave(url, &wg)

	}

	fmt.Printf("No. of Go Routines: %d\n", runtime.NumGoroutine())

	wg.Wait()
}
