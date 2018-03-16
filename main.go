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
	ssh = flag.String("ssh", "", "-ssh ssh port")
	port = flag.Int("port", 7000, "-port target ssh port")
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
			log.Printf("can not connect to ssh = %s, error=%s", *ssh, err.Error())
			continue
		}
		go func() {
			defer func() {
				recover()
			}()
			_, err := io.Copy(conn, sshCon)
			if err != nil {
				println(err)
				sshCon.Close()
				conn.Close()
				return
			}
			_, err = io.Copy(sshCon, conn)
			if err != nil {
				println(err)
				sshCon.Close()
				conn.Close()
				return
			}
		}()
	}
}

func getSSHConection() (net.Conn, error) {
	return net.Dial("tcp", *ssh)
}
