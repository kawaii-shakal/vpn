package main

import (
	"crypto/tls"
	"log"
	"net"

	"github.com/songgao/water"
)

func handleTunTraffic(ifce *water.Interface, serverConn net.Conn) {
	buffer := make([]byte, 1500)
	for {
		n, err := ifce.Read(buffer)
		if err != nil {
			log.Fatal("Error reading from TUN interface:", err)
		}

		obfuscatedPacket := obfuscateTraffic(buffer[:n])
		_, err = serverConn.Write(obfuscatedPacket)
		if err != nil {
			log.Fatal("Error sending data to server:", err)
		}
	}
}

func connectToServer(certFile, keyFile, caCertFile string) (net.Conn, error) {
	tlsConfig, err := loadTLSConfig(certFile, keyFile, caCertFile)
	if err != nil {
		return nil, err
	}

	conn, err := tls.Dial("tcp", "46.22.211.37:443", tlsConfig)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to server")
	return conn, nil
}
