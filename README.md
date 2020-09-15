# Sokoban

Implementation of the classic [Sokoban](https://en.wikipedia.org/wiki/Sokoban) in Go, using [Pixel](https://github.com/faiface/pixel) 2D Game library.

<p>
  <img src="https://raw.githubusercontent.com/csixteen/sokoban/master/screenshots/sokoban.png" />
</p>

# Dependencies

The project uses Go modules, so you'll need a version of Go more recent than [1.11](https://blog.golang.org/using-go-modules).

# Building and Running

```
$ make bin
go build -o soko cmd/sokoban/*.go
$ ./soko
```

# Controls

- Arrows Up, Down, Left and Right - move the character to the adjacent cell
- `q` - quits the game
- `r` - resets the level

# Board elements

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

# Adding a new level

The levels are defined in the `levels.dat` file. Each level is specified as 8 consecutive strings with 8 characters each. Different levels are separated by one or more empty strings. Right now each level must be a matrix 8x8, but this is a limitation that I'll hopefully fix soon.

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

# Limitations and caveats

- Levels are limited to being defined as 8x8 matrices.
- Hardcoded values (sprite sheets, sprites dimensions, etc).

# Contributing

There is still lots of room for improving things. Pull-requests with bug fixes or improvements are more than welcome.

# References

- Pixel 2D Game library: https://github.com/faiface/pixel
- SpriteSheets: https://kenney.nl/assets/sokoban

# License

[MIT](https://github.com/csixteen/sokoban/blob/master/LICENSE)
