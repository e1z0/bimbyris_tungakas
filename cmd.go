package main

import (
        "fmt"
	"flag"
	"os"
)

func Help() {
fmt.Printf("Dummy text\n")
os.Exit(0)
}

func CmdLine() {
        help := flag.Bool("help", false, "Help")
        flag.Parse()
        if *help {
            Help()
        }
}
