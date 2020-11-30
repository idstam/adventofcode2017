package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

func initScreen() tcell.Screen {
	// init screen
	encoding.Register()
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, err := tcell.NewScreen()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	//s.EnableMouse()
	s.Clear()
	return s
}
