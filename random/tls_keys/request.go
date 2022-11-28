package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "makes a GET HTTP request to the given URL")
	flag.Parse()

	if url == "" {
		fmt.Println("[Error] The URL cannot be empty")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := makeRequest(url); err != nil {
		log.Fatalln(err)
	}

}

func makeRequest(url string) error {
	f, err := os.OpenFile("/tmp/keys", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				KeyLogWriter: f,
			},
		},
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(key, "=", value)
		}
	}
	fmt.Println()
	fmt.Println()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes[:250]))

	return nil
}
