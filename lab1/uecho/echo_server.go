// Leave an empty line above this comment.
package main

import (
	"log"
	"net"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// UDPServer implements the UDP Echo Server specification found at
// lab1/uecho/README.md#udp-echo-server
type UDPServer struct {
	conn *net.UDPConn
}

// NewUDPServer returns a new UDPServer listening on addr. It should return an
// error if there was any problem resolving or listening on the provided addr.
func NewUDPServer(addr string) (*UDPServer, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}
	return &UDPServer{conn: conn}, nil
}

// ServeUDP starts the UDP server's read loop. The server should read from its
// listening socket and handle incoming client requests as according to the
// the specification.
func (u *UDPServer) ServeUDP() {
	for {
		buf := make([]byte, 1500)
		n, addr, err := u.conn.ReadFromUDP(buf)
		if err != nil {
			if socketIsClosed(err) {
				break
			}
			log.Printf("error reading from UDP connection: %v", err)
			continue
		}
		resp := handleCmd(buf[:n])
		_, err = u.conn.WriteToUDP([]byte(resp), addr)
		if err != nil {
			log.Printf("error writing to UDP connection: %v", err)
			continue
		}
	}
}

func handleCmd(b []byte) string {
	parts := strings.Split(string(b), "|:|")
	if len(parts) != 2 {
		return "Unknown command"
	}
	cmd, args := parts[0], parts[1]

	// cmd, args, ok := strings.Cut(string(b), "|:|")
	// if !ok {
	// 	return "Unknown command"
	// }
	// if cmd != "" && strings.Contains(args, "|:|") {
	// 	return "Unknown command"
	// }

	switch cmd {
	case "UPPER":
		return strings.ToUpper(args)
	case "LOWER":
		return strings.ToLower(args)
	case "CAMEL":
		return cases.Title(language.English).String(args)
	case "ROT13":
		rot13 := func(r rune) rune {
			switch {
			case 'A' <= r && r <= 'Z':
				return 'A' + (r-'A'+13)%26
			case 'a' <= r && r <= 'z':
				return 'a' + (r-'a'+13)%26
			}
			return r
		}
		return strings.Map(rot13, args)
	case "SWAP":
		return strings.Map(func(r rune) rune {
			switch {
			case unicode.IsLower(r):
				return unicode.ToUpper(r)
			case unicode.IsUpper(r):
				return unicode.ToLower(r)
			}
			return r
		}, args)
	default:
		return "Unknown command"
	}
}

// socketIsClosed is a helper method to check if a listening socket has been closed.
func socketIsClosed(err error) bool {
	return strings.Contains(err.Error(), "use of closed network connection")
}
