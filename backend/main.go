package main

import (
	"flag"
	"log"

	"github.com/rnd00/noet/backend/router"
)

type flags struct {
	Debug bool
}

func parseflag() *flags {
	debug := flag.Bool("debug", false, "debug flag, if true then info will be printed out")

	flag.Parse()

	return &flags{
		Debug: *debug,
	}
}

func main() {
	pf := parseflag()
	r := router.NewRouter()
	r.SetDebug(pf.Debug)
	r.Invoke()
	log.Fatal(r.Run())
}
