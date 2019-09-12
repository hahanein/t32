t32 - tic tac toe 2.0


How to play
===========

When it is your turn enter coordinates in the form of NxN where N is a natural
number. To place your player in the first square in the top-left corner enter
0x0.


Build
=====

Download and install [Go >1.11][1]. Then build and run the binary with the
following commands on Linux:

        $ go build -o bin ./cmd/
        $ ./bin

On Windows:

        $ go.exe build -o bin.exe .\cmd\
        $ .\bin.exe


Configuration
=============

Place a config.json in the current directory or use its path as an argument to
the -config flag on Linux:

        $ ./bin -config /path/to/config.json

On Windows:
        
        $ .\bin.exe -config .\path\to\config.json

The configuration file must be formatted as follows:

        {
          "size": 3,
          "player1": "A",
          "player2": "B",
          "player3": "C"
        }



Architecture
============

This program is made up of several layers:

                      ,__________________________________,
                      | DRIVERS AND FRAMEWORKS           |
                      | ======================           |
                      |                                  |
                      | - t32/ai                         |
                      | ,______________________________, |
                      | | INTERFACE ADAPTERS           | |
                      | | ==================           | |
                      | |                              | |
                      | | - t32/clients/computer       | |
                      | | - t32/clients/console        | |
                      | | ,__________________________, | |
                      | | | USE CASES                | | |
                      | | | =========                | | |
                      | | |                          | | |
                      | | | - t32/actors/referee     | | |
                      | | | - t32/actors/participant | | |
                      | | |      ,____________,      | | |
                      | | |      | ENTITIES   |      | | |
                      | | |      | ========   |      | | |
                      | | |      |            |      | | |
                      | | |      | - t32/game |      | | |
                      | | |      '------------'      | | |
                      | | '--------------------------' | |
                      | '------------------------------' |
                      '----------------------------------' 

The outer circles are mechanisms. The inner circles are policies. Source code
dependencies may only point inwards. This produces:

- independence of any external agency such as a framework, user interface or
  database
- testability

Of course rules exist to be broken and so game/json.go encapsulates knowledge
about the outside world (encoding formats) but it is merely a few lines of
boiler plate that have zero effect on the rest of that package and the
best alternative might be much more confusing.

Participants react asynchronously to changes in the state of the game. This
opens up the possibilities for alternative, distributed implementations of the
Client interface that are more common in modern multiplayer games.


Packages
========

t32/game
--------

Holds the Game which provides methods to mutate itself in a limited number of
ways. Everything you may or may not be allowed to do in a Game is encapsulated
in this package.


t32/actors
----------

Holds the Referee which serves the double purpose of (1) managing the Game's
single source of truth and (2) publishing its state. It also holds the
Participant which provides a unified interface through which external agencies
may participate in the Game either as Players or as spectators.


t32/clients
-----------

Holds adapters which connect external agencies like a human user or an
artificial intelligence to a Participant.

You may use godoc to read the documentation. For example you may read the docs
for the game package on the command line with the following command on Linux:

        $ godoc ./game/ | less


Run Tests
=========

Use the following command to run all available tests on Linux:

        $ go test ./...

On Windows:

        $ go.exe test .\...


---

[1]: https://golang.org/dl/
