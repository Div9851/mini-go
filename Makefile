.PHONY: test clean

BINDIR		= bin
BINARY		= mgo
GO_FILES	= $(shell find . -type f -name '*.go')

$(BINARY): $(GO_FILES)
	go build -o $(BINDIR)/$(BINARY) cmd/main.go

test: $(BINARY)
	./test.sh

clean:
	rm $(BINDIR)/$(BINARY) tmp*
