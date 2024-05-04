.DEFAULT_GOAL=prod

GO=go
PRODARGS=-ldflags "-s -w"
DEBUGARGS=debug

run:
	$(GO) run . $(DEBUGARGS)

run-prod:
	$(GO) run .

build:
	$(GO) build

prod:
	$(GO) build $(PRODARGS)

clean:
	$(GO) clean

.PHONY: run build prod clean
