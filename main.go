package main

import (
	"flag"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	ofm, tls                       bool
	port, bindAddr, dir, cert, key string
)

func main() {
	setFlags()
	route := gin.New()
	setupTree(route, dir)
	if tls {
		route.RunTLS(bindAddr+":"+port, cert, key)
	} else {
		route.Run(bindAddr + ":" + port)
	}
}
func setFlags() {
	flag.BoolVar(&ofm, "ofm", false, "")
	flag.StringVar(&port, "port", "8000", "")
	flag.StringVar(&bindAddr, "addres", "127.0.0.1", "")
	flag.StringVar(&cert, "cert", "", "")
	flag.StringVar(&key, "key", "", "")
	flag.Parse()
	if ((cert != "") && (key == "")) || ((cert == "") && (key != "")) {
		log.Fatal("You must set cert of key file.\nUse -cert and -key for this.")
	}
	if cert != "" && key != "" {
		tls = true
	}
	_, err := net.ResolveTCPAddr("tcp", bindAddr+":"+port)
	if err != nil {
		log.Fatalf("Could not resolve the address to listen to: %s", bindAddr+":"+port)
	}
	dir = os.Args[len(os.Args)-1]

}

func setupTree(r *gin.Engine, folder string) {
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			r.GET(path, func(c *gin.Context) {
				c.File(path)
			})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
