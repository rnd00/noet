package main

import (
	"flag"
	"log"
	"time"

	"github.com/rnd00/noet/backend/controllers"
	"github.com/rnd00/noet/backend/router"
)

type flags struct {
	Debug bool
	Port  int
}

func parseflag() *flags {
	debug := flag.Bool("debug", false, "debug flag, if true then info will be printed out")
	port := flag.Int("port", 8050, "port number, if empty then set to 8050")

	flag.Parse()

	return &flags{
		Debug: *debug,
		Port:  *port,
	}
}

// router-gin
// func main() {
// 	pf := parseflag()
// 	r := router.NewRouter()
// 	r.SetDebug(pf.Debug)
// 	r.Invoke()
// 	log.Fatal(r.Run())
// }

// testing router-pure
func main() {
	pf := parseflag()

	h := router.NewHandler()
	h.SetupMuxer(router.GET, "/testwrite", controllers.TestWrite)
	httpHandler := h.ReturnHttpHandler()

	r := router.NewRoutern()
	r.SetHandler(&httpHandler)
	r.SetPort(pf.Port)
	r.SetTimeout(2 * time.Second)
	err := r.Invoke()
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(r.Run())
}
