package main

import (
	"fmt"
	"time"
	// "bytes"
	"github.com/foize/go.sgr"
	"os"
	"os/signal"
)

var (
	colors = [...]uint8{196, 214, 226, 154, 46, 49, 51, 39, 21, 129, 201, 199}
	gw = "Lᴇɢᴀʟɪᴢᴇ Gᴀʏ Wᴇᴇᴅ"
)

func legalizer() func() string {
	pad  := 0
	return func() string {
		if pad > len(colors) { pad -= len(colors) }
		var buf = ""
		//var buf = bytes.Buffer
		for i,c := range gw { buf+=(sgr.FgColor(colors[(i+pad) % len(colors)])+string(c)) }
		// for i,c := range gw { buf.WriteString(sgr.FgColor(colors[(i+pad) % len(colors)])+string(c)) }
		pad += 1;
		return buf
		//return buf.String()
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for c := range c {
			if c == os.Interrupt {
				fmt.Print("\x1b[?25h\n")
				os.Exit(0)
			}
		}
	}()

	lgw := legalizer()
	fmt.Print("\x1b[?25l")
	for {
		fmt.Print("\r\033[0K"+lgw())
		time.Sleep(100 * time.Millisecond)
	}
}
