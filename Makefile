.DEFAULT_GOAL=prod

GO=go
PRODARGS=-ldflags "-s -w"
DEBUGARGS=debug

run:
	$(GO) run . $(DEBUGARGS)

build:
	$(GO) build

prod:
	$(GO) build $(PRODARGS)

clean:
	$(GO) clean

.PHONY: run build prod clean
