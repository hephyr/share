package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
	"github.com/iozephyr/share/serve"
)

var dir string
var port int

func init() {
	flag.StringVar(&dir, "d", "", "set share directory")
	flag.IntVar(&port, "p", 8080, "set server port")
	flag.Parse()
}

func main() {
	if dir == "" {
		dir = flag.Arg(0)
	}
	if dir != "" {
		if !filepath.IsAbs(dir) {
			currentDir, _ := os.Getwd()
			dir = filepath.Join(currentDir, dir)
		}
	}
	http.Handle("/", http.RedirectHandler("/files/", http.StatusMovedPermanently))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(rice.MustFindBox("static").HTTPBox())))
	http.Handle("/files/", http.StripPrefix("/files/", serve.FileServer(serve.Dir(dir))))

	ipaddr := getIP()
	fmt.Printf("Server address: http://%s:%d\n", ipaddr, port)

	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, nil)
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
