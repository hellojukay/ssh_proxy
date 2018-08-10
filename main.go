package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var port *int
var ssh *string

func init() {
	ssh = flag.String("target", "", "-target target server and port: 23.65.123.3:90")
	port = flag.Int("listen", 7000, "-listen listen a local port")
	flag.Parse()
}
func main() {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Printf("bind tcp ,port=%d ,error, %s\n", *port, err.Error())
		os.Exit(1)
	}
	log.Printf("bind success on port %d", *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("handle connection error=%s", err.Error())
			continue
		}
		sshCon, err := getSSHConection()
		if err != nil {
			log.Printf("can not connect to target,target = %s, error=%s", *ssh, err.Error())
			continue
		}
		go func() {
			io.Copy(conn, sshCon)
		}()
		go func() {
			io.Copy(sshCon, conn)
		}()
	}
}

func getSSHConection() (net.Conn, error) {
	return net.Dial("tcp", *ssh)
}
