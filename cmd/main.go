package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"t32/actors/participant"
	"t32/actors/referee"
	"t32/ai"
	"t32/clients/computer"
	"t32/clients/console"
	"t32/clients/console/templates"
	"t32/game"
)

var (
	version = "v0.1.0"

	// Comand line flags.
	pathToConfig   string
	displayVersion bool
)

func init() {
	flag.StringVar(&pathToConfig, "config", "config.json", "path to config")
	flag.BoolVar(&displayVersion, "version", false, "display version and stop execution")

	// // Log to syslog instead of printing to stdout.
	// w, err := syslog.New(syslog.LOG_ERR, "t32")
	// if err != nil {
	// 	panic(err)
	// }
	// log.SetOutput(w)
}

func main() {
	flag.Parse()

	if displayVersion {
		fmt.Printf("t32 %s\n", version)

		return
	}

	cfg, err := configFromFilepath(pathToConfig)
	if err != nil {
		log.Fatal(err)
	}

	g, err := game.New(cfg.Size)
	if err != nil {
		log.Fatal(err)
	}

	r := referee.New(*g)

	cConsole := console.New(
		new(templates.Templates),
		new(console.Output),
		os.Stdin,
	)

	cComputer := computer.New(ai.Random)

	p1 := participant.New(cfg.Player1, cConsole, r)
	p2 := participant.New(cfg.Player2, cConsole, r)
	p3 := participant.New(cfg.Player3, cComputer, r)

	<-p1.Done
	<-p2.Done
	<-p3.Done
}