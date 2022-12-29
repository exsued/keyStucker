package main

import (
	"github.com/karalabe/usb"
	"log"
)

func main() {
	devices, err := usb.Enumerate(1578, 19457)
	if err != nil {
		log.Println(err)
	}else{
		device, err := devices[0].Open()
		if err != nil {
			log.Println(err)
		}
		data := make([]byte, 256)
		for {
			_, err := device.Read(data)
			log.Println(data)
			if err != nil {
				log.Println(err)
			}
		}
		device.Close()
	}
}
