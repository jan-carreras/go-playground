package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {

	if err := run(); err != nil {
		log.Fatalln(err)
	}

}

func run() error {
	start := time.Now()
	resp, err := http.Get("https://example.org")
	fmt.Printf("HTTP Call time: %v\n\n", time.Now().Sub(start))

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("HEADERS")
	for k, v := range resp.Header {
		fmt.Printf("\t%s: %v\n", k, v)
	}

	fmt.Println("BODY")
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(r)[:150])

	return nil
}
