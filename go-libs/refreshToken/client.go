package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

type Client struct {
	token        Token
	server       *Server
	singleflight *singleflight.Group
}

func NewClient(server *Server) *Client {
	return &Client{server: server, singleflight: &singleflight.Group{}}
}

func (c *Client) Do() {
	retries := 3
	for i := 0; i < retries; i++ {
		if c.token.IsExpired() {
			if err := c.refreshToken(); err != nil {
				fmt.Println("[client] Unable to refresh the token", err)
			}
		}

		err := c.server.DoSomething(c.token, "do some action")
		if errors.Is(err, errTokenExpired) {
			fmt.Println("[Client] Retrying:", err)
			break
		}

		if err != nil {
			fmt.Println("[Client] Error from server:", err)
		}
	}

}

func (c *Client) refreshToken() error {
	token, err, _ := c.singleflight.Do("token-refresh", func() (interface{}, error) {
		token, err := c.server.Refresh()
		return token, err
	})

	fmt.Println("[client] Refreshing the token", c.token.key)
	if err != nil {
		return fmt.Errorf("server.Refresh: %w", err)
	}
	c.token = token.(Token)

	return nil
}

func main() {
	maxRequest := 3 * time.Second
	tokenTTL := 5 * time.Second

	server := NewServer(maxRequest, tokenTTL)
	client := NewClient(server)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(999 * time.Millisecond)
			defer wg.Done()
			for {
				client.Do()
			}
		}()

	}

	wg.Wait()
}
