package main

import (
	"flag"
	"fmt"
	"github.com/tarm/serial"
	"log"
)

type serialReader struct {
	name   string
	serial *serial.Port
}

func main() {
	c1 := flag.String("c1", "/dev/ttyUSB0", "serial 1 name")
	c2 := flag.String("c2", "/dev/ttyUSB1", "serial 2 name")
	baud := flag.Int("baud", 115200, "baud rate")
	flag.Parse()

	s1, err := newSerialReader(*c1, *baud)
	if err != nil {
		log.Fatal(err)
	}

	s2, err := newSerialReader(*c2, *baud)
	if err != nil {
		log.Fatal(err)
	}

	go s1.readLoop()
	s2.readLoop()
}

func newSerialReader(name string, baud int) (*serialReader, error) {
	s, err := serial.OpenPort(&serial.Config{Name: name, Baud: baud})
	if err != nil {
		return nil, err
	}

	return &serialReader{
		name:   name,
		serial: s,
	}, nil
}
func (s serialReader) readLoop() {
	for {
		b := make([]byte, 256)
		_, err := s.serial.Read(b)
		if err != nil {
			fmt.Printf("Read error %s, error: %s\n", s.name, err.Error())
		} else {
			fmt.Printf("got %s: error: %s", s.name, string(b))
		}
	}
}
