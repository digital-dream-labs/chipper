.PHONY: build

build: 
	CGO_ENABLED=0 go build \
	-ldflags "-w -s -extldflags "-static"" \
	-trimpath \
	-o chipper cmd/main.go

build-linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 build \
	-ldflags "-w -s -extldflags "-static"" \
	-trimpath \
	-o chipper cmd/main.go