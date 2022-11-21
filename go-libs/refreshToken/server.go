package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var errTooManyRequests = errors.New("too many requests")
var errTokenExpired = errors.New("expired token")
var errInvalidToken = errors.New("invalid Token")

type Token struct {
	key       string
	expiresAt time.Time
}

func (t Token) IsExpired() bool {
	return time.Now().Sub(t.expiresAt) >= 0
}

func (t Token) Invalid(otherToken Token) bool {
	return t.key != otherToken.key
}

type Server struct {
	lastTime   time.Time
	maxRequest time.Duration
	tokenTTL   time.Duration
	validToken Token
}

func NewServer(maxRequest time.Duration, tokenTTL time.Duration) *Server {
	return &Server{maxRequest: maxRequest, tokenTTL: tokenTTL}
}

func (s *Server) DoSomething(token Token, request string) error {
	if s.validToken.Invalid(token) {
		//fmt.Println("[server] Invalid token")
		return errInvalidToken
	}

	if s.validToken.IsExpired() {
		//fmt.Println("[server] Token expired")

		return errTokenExpired
	}

	//fmt.Println("[server] Token is valid")

	//fmt.Printf("[server] Doing request: %s\n", request)

	return nil
}

func (s *Server) Refresh() (Token, error) {
	//fmt.Println("[server] Refresh called")
	if !s.lastTime.IsZero() && time.Now().Sub(s.lastTime) < s.maxRequest {
		//fmt.Println("[server] Too many requests")
		return Token{}, errTooManyRequests
	}

	time.Sleep(100 * time.Millisecond)

	//fmt.Println("[server] Token refreshed")
	s.lastTime = time.Now()
	s.validToken = Token{
		key:       s.newToken(),
		expiresAt: time.Now().Add(s.tokenTTL),
	}
	return s.validToken, nil

}

func (s *Server) newToken() string {
	return fmt.Sprintf("token-%d", rand.Int())
}
