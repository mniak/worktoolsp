package main

import (
	"github.com/mniak/hsmlib"
	_ "github.com/mniak/hsmlib"
	_ "github.com/mniak/krypton"
)

func main() {
	packet := hsmlib.Packet{
		Header:  RandomHeader(),
		Payload: payload,
	}
	return conn.SendPacket(packet)
}
