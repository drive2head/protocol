package main

import (
	"fmt"
	"net"
)

type SessionStatus string

const (
	ACTIVE SessionStatus = "ACTIVE"
	CLOSED SessionStatus = "CLOSED"
)

type Session struct {
	Addr   string
	Status SessionStatus
}

var sessions map[string]Session

func OpenSession(c net.Conn) {
	addr := c.RemoteAddr().String()
	sessions[addr] = Session{
		Addr:   addr,
		Status: ACTIVE,
	}

	fmt.Printf("Opened new session for %s\n", addr)
}

func CloseSession(c net.Conn) {
	addr := c.RemoteAddr().String()
	sessions[addr] = Session{
		Addr:   addr,
		Status: CLOSED,
	}

	fmt.Printf("Session closed for %s\n", addr)
}
