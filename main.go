package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var (
	fs          = flag.NewFlagSet("hacker-typing", flag.ExitOnError)
	speed       = fs.Int("speed", 1, "characters per keypress")
	autotypeMin = fs.Duration("autotype-min", 0, "autotype min duration")
	autotypeMax = fs.Duration("autotype-max", 0, "autotype max duration (suggested value: 200ms)")
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
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

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

	// writing routine
	writeCh := make(chan int, 0)
	go func() {
		pos := 0
		for {
			select {
			case n := <-writeCh:
				for n > 0 {
					if pos >= len(content) { // loop
						pos = 0
					}
					fmt.Printf("%c", content[pos])
					pos++
					n--
				}
			case <-ctx.Done():
				fmt.Println() // newline
				return
			}
		}
	}()

	// autotype
	if *autotypeMin >= 0 && *autotypeMax >= *autotypeMin {
		go func() {
			writeCh <- 1 // immediately type a char
			for {
				duration := float64(*autotypeMin) + rand.Float64()*(float64(*autotypeMax)-float64(*autotypeMin))
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Duration(duration)):
					writeCh <- 1
				}
			}
		}()
	}

	// init keyboard
	{
		if err := keyboard.Open(); err != nil {
			return err
		}
		defer keyboard.Close()
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				return err
			}
			switch key {
			case keyboard.KeyEsc, keyboard.KeyCtrlC:
				cancel()
				return nil
			default:
				writeCh <- *speed
			}
		}
	}
}
