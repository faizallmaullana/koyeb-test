package main

import (
	"fmt"
	"net"
)

func getIP() {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, addr := range addr {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// check if IPv4 or IPv6 is not nil
			if ipnet.IP.To4() != nil || ipnet.IP.To16 != nil {
				// print available addresses
				fmt.Print("Server is running on: ")
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
