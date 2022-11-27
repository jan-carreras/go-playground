package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	fmt.Println("Establishing the TLS connection...")
	conn, err := tls.Dial("tcp", "example.org:443", &tls.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Established.")
	defer conn.Close()

	fmt.Println("Marking the HTTP request...")
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: example.org\r\nConnection: close\r\n\r\n"))
	if err != nil {
		return err
	}
	fmt.Println("Done.")

	fmt.Println("Reading HTTP response...")
	buf, err := io.ReadAll(conn)
	if err != nil {
		return err
	}
	fmt.Println("Done.")

	response := string(buf)

	headerIdx := strings.Index(response, "\r\n")

	fmt.Println("HTTP Header response:")
	fmt.Println("\t", response[:headerIdx])

	endHeadersIndex := strings.Index(response, "\r\n\r\n")
	fmt.Println("HTTP Headers response:")

	for _, header := range strings.Split(response[headerIdx+2:endHeadersIndex], "\r\n") {
		fmt.Println("\t", header)
	}

	body := response[endHeadersIndex+4:]
	fmt.Println("HTTP Header response:")
	fmt.Println(body[:150])

	return nil
}
