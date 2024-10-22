// server.go
package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"

	"github.com/songgao/water"
)

func handleClient(conn net.Conn, ifce *water.Interface) {
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		n, err := ifce.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client")
				return
			}
			log.Println("Error reading from TUN interface:", err)
			continue
		}

		obfuscatedPacket := obfuscateTraffic(buffer[:n])
		_, err = conn.Write(obfuscatedPacket)
		if err != nil {
			log.Println("Error sending data to client:", err)
			return
		}
	}
}

func startTLSServer(certFile, keyFile, caCertFile string, ifce *water.Interface) {
	tlsConfig, err := loadTLSConfig(certFile, keyFile, caCertFile)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("TLS VPN Server is listening on port 443")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		go handleClient(conn, ifce)
	}
}
