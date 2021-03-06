.PHONY: test
test:
	go test -v pkg/utils/*.go
	go test -v pkg/game/*.go

.PHONY: bin
bin:
	pkger -o cmd/sokoban/
	go build -o soko cmd/sokoban/*.go
