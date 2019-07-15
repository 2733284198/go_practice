package main

import (
	"fmt"
	"log"
	"os"

	"gosrc.io/xmpp"
)

func main() {
	config := xmpp.Config{
		Address:      "172.17.20.39:5222",
		Jid:          "frank03___hst@xxxx",
		Password:     "1",
		PacketLogger: os.Stdout,
		Insecure:     false,
	}

	client, err := xmpp.NewClient(config)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// If you pass the client to a connection manager, it will handle the reconnect policy
	// for you automatically.
	cm := xmpp.NewStreamManager(client, nil)
	err = cm.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Iterator to receive packets coming from our XMPP connection
	for packet := range client.Recv() {
		switch packet := packet.(type) {
		case xmpp.Message:
			_, _ = fmt.Fprintf(os.Stdout, "Body = %s - from = %s\n", packet.Body, packet.From)
			reply := xmpp.Message{PacketAttrs: xmpp.PacketAttrs{To: packet.From}, Body: packet.Body}
			_ = client.Send(reply)
		default:
			_, _ = fmt.Fprintf(os.Stdout, "Ignoring packet: %T\n", packet)
		}
	}
}