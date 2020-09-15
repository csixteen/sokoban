# Sokoban

Implementation of the classic [Sokoban](https://en.wikipedia.org/wiki/Sokoban) in Go.

<p>
  <img src="https://raw.githubusercontent.com/csixteen/sokoban/master/screenshots/sokoban.png" />
</p>

# Building and Running

```
$ make bin
go build -o soko cmd/sokoban/*.go
$ ./soko
```

# Controls

- Arrows Up, Down, Left and Right - move the character to the adjacent cell

# Testing

```
$ make test
go test -v pkg/utils/*.go
=== RUN   TestPushToStack
--- PASS: TestPushToStack (0.00s)
=== RUN   TestPopFromStack
--- PASS: TestPopFromStack (0.00s)
=== RUN   TestPopFromEmptyStack
--- PASS: TestPopFromEmptyStack (0.00s)
=== RUN   TestTopFromEmptyStack
--- PASS: TestTopFromEmptyStack (0.00s)
PASS
ok  	command-line-arguments	(cached)
go test -v pkg/types/*.go
=== RUN   TestNewBoard
--- PASS: TestNewBoard (0.00s)
=== RUN   TestMovePlayer
--- PASS: TestMovePlayer (0.00s)
=== RUN   TestMovePlayerTwoBlocks
--- PASS: TestMovePlayerTwoBlocks (0.00s)
PASS
ok  	command-line-arguments	(cached)
```

# References
- SpriteSheets: https://kenney.nl/assets/sokoban
