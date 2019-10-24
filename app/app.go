package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/carlmjohnson/flagext"
	"github.com/peterbourgon/ff"
	"github.com/spotlightpa/arc-reader/internal/feed"
)

const AppName = "arc-reader"

func CLI(args []string) error {
	a, err := parseArgs(args)
	if err != nil {
		return err
	}
	if err := a.exec(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return err
	}

	return nil
}

func parseArgs(args []string) (*app, error) {
	fl := flag.NewFlagSet(AppName, flag.ContinueOnError)
	src := flagext.FileOrURL(flagext.StdIO, nil)
	fl.Var(src, "src", "source file or URL")
	l := log.New(nil, AppName+" ", log.LstdFlags)
	fl.Var(
		flagext.Logger(l, flagext.LogVerbose),
		"verbose",
		`log debug output`,
	)

	fl.Usage = func() {
		fmt.Fprintf(fl.Output(), `arc-reader

Usage:

	arc-reader [options]

Options:
`)
		fl.PrintDefaults()
	}
	if err := ff.Parse(fl, args, ff.WithEnvVarPrefix("ARC_READER")); err != nil {
		return nil, err
	}

	id := fl.Arg(0)
	a := app{src, l, id}
	return &a, nil
}

type app struct {
	src io.ReadCloser
	*log.Logger
	id string
}

func (a *app) exec() (err error) {
	a.Println("starting")
	dec := json.NewDecoder(a.src)
	var f feed.Feed
	if err = dec.Decode(&f); err != nil {
		return err
	}
	var body string
	var v interface{} = f
	if a.id != "" {
		v = nil
		for _, s := range f.Stories {
			if s.ID == a.id {
				v = s
				body = s.Body
			}
		}
	}

	io.WriteString(os.Stdout, "+++\n")
	enc := toml.NewEncoder(os.Stdout)
	if err = enc.Encode(v); err != nil {
		return err
	}
	io.WriteString(os.Stdout, "+++\n\n"+body+"\n")
	return err
}
