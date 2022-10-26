package main

import (
	"fmt"
	"github.com/armon/go-socks5"
	"os"
)

var (
	listenAddr = "0.0.0.0:1080"
	username   = ""
	password   = ""
	conf       = &socks5.Config{}
)

func init() {
	switch len(os.Args) {
	case 1:
		//	socks5://0.0.0.0:1080
	case 2:
		//	socks5://listenAddr
		listenAddr = os.Args[1]
	case 3:
		//	socks5://username@listenAddr
		listenAddr = os.Args[1]
		username = os.Args[2]
	case 4:
		//	socks5:/username:password@listenAddr
		listenAddr = os.Args[1]
		username = os.Args[2]
		password = os.Args[3]
	default:
		fmt.Println("socks5d [listenAddr] [username] [password]")
		os.Exit(0)
	}
	if username != "" {
		cred := socks5.StaticCredentials{
			username: password,
		}
		cator := socks5.UserPassAuthenticator{Credentials: cred}
		conf.AuthMethods = append(conf.AuthMethods, cator)
	}
}

func main() {
	server, err := socks5.New(conf)
	if err != nil {
		fmt.Println("[error]", err)
		os.Exit(0)
	}
	// Create SOCKS5 proxy on localhost port 1080
	fmt.Printf("listening socks5://%s\n", listenAddr)

	if err := server.ListenAndServe("tcp", listenAddr); err != nil {
		fmt.Println("[error]", err)
		os.Exit(0)
	}
}
