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

	// Create a new Game.
	g, err := game.New(cfg.Size)
	if err != nil {
		log.Fatal(err)
	}

	// Pass that Game to a new Referee.
	r := referee.New(*g)

	// Create a multiplayer console client.
	cConsole := console.New(
		templates.New(),
		new(console.Output),
		os.Stdin,
	)

	// Create an automated computer client.
	cComputer := computer.New(ai.Random)

	// Create the participants and automatically register them with the
	// Referee. They will immediately begin interacting with the Game.
	p1 := participant.New(cfg.Player1, cConsole, r)
	p2 := participant.New(cfg.Player2, cConsole, r)
	p3 := participant.New(cfg.Player3, cComputer, r)

	// Block until all Participants declared them being done by having send
	// an empty struct through the Done channel.
	<-p1.Done
	<-p2.Done
	<-p3.Done
}
