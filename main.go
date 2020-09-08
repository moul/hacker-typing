package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

var speed = flag.Int("speed", 1, "characters per keypress")

func main() {
	flag.Parse()
	args := flag.Args()

	log.SetFlags(0)
	if len(args) < 1 {
		log.Print("usage: hacker-typing [flags] path/to/file")
		os.Exit(1)
	}

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	i := 0
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		for j := 0; j < *speed; j++ {
			if i < len(content) {
				fmt.Printf("%c", content[i])
			}
			i++
		}
	}
}
