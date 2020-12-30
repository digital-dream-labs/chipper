.PHONY: build

build: 
	CGO_ENABLED=0 go build \
	-ldflags "-w -s -extldflags "-static"" \
	-trimpath \
	-o chipper cmd/main.go