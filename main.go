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
var ssh *int

func init() {
	ssh = flag.Int("ssh", 22, "-ssh ssh port")
	port = flag.Int("p", 7000, "-port target ssh port")
	flag.Parsed()
}
func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Printf("bind tcp ,port=%d ,error, %s\n", *port, err.Error())
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("handle connection error=%s", err.Error())
			continue
		}
		sshCon, err := getSSHConection()
		if err != nil {
			log.Printf("can not connect to ssh , port=%d, error=%s", *ssh, err.Error())
			continue
		}
		go io.Copy(conn, sshCon)
		go io.Copy(sshCon, conn)
	}
}

func getSSHConection() (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", ssh))
}
