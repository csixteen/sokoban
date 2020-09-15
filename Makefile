.PHONY: test
test:
	go test -v pkg/utils/*.go
	go test -v pkg/types/*.go

.PHONY: bin
bin:
	go build -o soko cmd/sokoban/*.go