package main

import "encoding/binary"

func obfuscateTraffic(packet []byte) []byte {
	if len(packet) < 20 {
		return packet
	}
	packet[8] = 64
	flagsOffset := binary.BigEndian.Uint16(packet[6:8])
	flagsOffset |= 0x2000
	binary.BigEndian.PutUint16(packet[6:8], flagsOffset)
	return packet
}
