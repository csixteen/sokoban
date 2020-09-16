# Sokoban

Implementation of the classic [Sokoban](https://en.wikipedia.org/wiki/Sokoban) in Go, using [Pixel](https://github.com/faiface/pixel) 2D Game library.

<p>
  <img src="https://raw.githubusercontent.com/csixteen/sokoban/master/screenshots/sokoban.png" />
</p>

# Installing

```
$ go get github.com/csixteen/sokoban/cmd/sokoban
```

# Building and Running

The project uses Go modules, so you'll need a version of Go more recent than [1.11](https://blog.golang.org/using-go-modules).
You will also need to have [pkger](https://github.com/markbates/pkger) installed.

```
$ make bin
go build -o soko cmd/sokoban/*.go
$ ./soko
```

# Controls

- Arrows Up, Down, Left and Right - move the character to the adjacent cell
- `q` - quits the game
- `r` - resets the level

# Adding a new level

The levels are defined in the `levels.dat` file. Each level is specified as a number of consecutive lines, all with the same length (essentially, a matrix).

The matrices in `levels.dat` will be rendered 90 degrees rotated anti-clockwise. I still haven't fixed this, but I don't think it's that big of a deal. Here is the matrix that corresponds to the level on the screenshot:

```
wwwwwwww
wffwfffw
wjbggbfw
wfbgbffw
wffggbfw
wffffffw
wfwwfwww
wwwwwwww
```

You'll have to run `make bin` after you make changes to **levels.dat**.

## Board elements

These are the current board elements:

- `w` - wall (unmovable)
- `b` - block (movable)
- `f` - floor
- `g` - goal (where you need to put the block onto)
- `o` - block on a goal (becomes unmovable)
- `h` - player facing left
- `j` - player facing down
- `k` - player facing up
- `l` - player facing right

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

# To Do

- Fix the orientation of the board. Right now, the level description in `levels.dat` results in a board that is rotated 90 degrees anti-clockwise when the window is rendered.
- Add more levels.
- Keep track of time that it takes to complete each level and display it.

# Contributing

There is still lots of room for improving things. Pull-requests with bug fixes or improvements are more than welcome.

# References

- Pixel 2D Game library: https://github.com/faiface/pixel
- SpriteSheets: https://kenney.nl/assets/sokoban
- Pkger: https://github.com/markbates/pkger. I'm using this tool to embedded the static assets into the binary. Thanks to [skuzzymiglet](https://www.reddit.com/r/golang/comments/itbr8t/im_still_fairly_new_to_go_so_i_decided_to/g5ghbrn/) for the suggestion.

# License

[MIT](https://github.com/csixteen/sokoban/blob/master/LICENSE)
