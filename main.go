package main

import (
	"flag"
	"log"
)

//func main() {
//	ifce, err := createTunInterface()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	serverConn, err := connectToServer("client-cert.pem", "client-key.pem", "ca-cert.pem")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	go handleTunTraffic(ifce, serverConn)
//
//	select {}
//}
//
//func main() {
//	ifce, err := createTunInterface()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	go func() {
//		startTLSServer("server-cert.pem", "server-key.pem", "ca-cert.pem", ifce)
//	}()
//
//	select {}
//}

func main() {
	runServer := flag.Bool("useServer", false, "Run as a server")
	flag.Parse()
	ifce, err := createTunInterface()
	if err != nil {
		log.Fatal(err)
	}
	if *runServer {
		go func() {
			startTLSServer("server-cert.pem", "server-key.pem", "ca-cert.pem", ifce)
		}()

		select {}
	} else {
		serverConn, err := connectToServer("client-cert.pem", "client-key.pem", "ca-cert.pem")
		if err != nil {
			log.Fatal(err)
		}

		go handleTunTraffic(ifce, serverConn)

		select {}
	}
}
