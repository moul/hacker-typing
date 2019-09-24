package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Print("usage: hacker-typing path/to/file")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
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
		if i < len(content) {
			fmt.Printf("%c", content[i])
		}
		i++
	}
}
