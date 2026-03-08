VERSION := $(shell git describe --tags --always --dirty)
BINARY := ship-htmx-app
MODULE := github.com/pnpancholi/ship-htmx-app

.PHONY : build run clean release

build:
	go build -ldflags "-X '${MODULE}/cmd.version=${VERSION}'" -o ${BINARY} .

run: 
	go run -ldflags "-X '${MODULE}/cmd.version=${VERSION}'" main.go 

clean: 
	rm -f ${BINARY}

release: 
	@read -p "Version (e.g 0.1.0): " v; \
	git tag v$$v; \
	git push origin main; \
	git push origin v$$v
