package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var (
	fs    = flag.NewFlagSet("hacker-typing", flag.ExitOnError)
	speed = fs.Int("speed", 1, "characters per keypress")
)

func main() {
	err := run(os.Args)
	if err != nil && err != flag.ErrHelp {
		log.Fatalf("error: %v", err)
	}

}

func run(osArgs []string) error {
	root := &ffcli.Command{
		ShortUsage: "hacker-typing [flags] <file>",
		ShortHelp:  "impress your friends",
		FlagSet:    fs,
		Exec:       hackerTypingSimulator,
	}
	return root.ParseAndRun(context.Background(), osArgs[1:])
}

func hackerTypingSimulator(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return flag.ErrHelp
	}

	// open source file
	var content []byte
	{
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()
		content, err = ioutil.ReadAll(f)
		if err != nil {
			return err
		}
	}

	// init keyboard
	{
		if err := keyboard.Open(); err != nil {
			return err
		}
		defer keyboard.Close()
	}
	i := 0
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			return err
		}
		switch key {
		case keyboard.KeyEsc, keyboard.KeyCtrlC:
			fmt.Println() // newline
			return nil
		default:
			for j := 0; j < *speed; j++ {
				if i < len(content) {
					fmt.Printf("%c", content[i])
				}
				i++
				// FIXME: add an option to quit or to loop at the end of the file
			}
		}
	}
}
