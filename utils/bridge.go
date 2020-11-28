package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
)

//Br ...
type Br struct {
	Connect    func(chan string)
	SendStanza func(string)
	Refresh    func()
}

//Tls client configurations
var config tls.Config = tls.Config{InsecureSkipVerify: true}
var c, _ = tls.Dial("tcp", "talk1110an.kik.com:5223", &config)

//Bridge ...
var Bridge = Br{
	Connect: func(msgr chan string) { //start server connection
		msg := make([]byte, 32768)
		for {
			n, err := c.Read(msg)
			if err != nil {
				fmt.Println(err)
				break
			}
			log.Printf("client: read %q (%d bytes)\n\n", string(msg[:n]), n)
			msgr <- string(msg[:n])
		}
		close(msgr)
	},
	SendStanza: func(stanza string) { //send stanza to the server
		n, err := io.WriteString(c, stanza)
		if err != nil {
			log.Fatalf("client: write: %s", err)
		}
		log.Printf("client: wrote %q (%d bytes)\n\n", stanza, n)
	},
	Refresh: func() { //refresh server connection to complete login
		c, _ = tls.Dial("tcp", "talk1110an.kik.com:5223", &config)
	},
}
