ALERTMANAGER_VERSION := $(shell ./alertmanager --version | cut -d" " -f3)

clean: 
	rm alertmanager

build: vet
	go build -tags netgo

docker-build: build
	 docker build -t alertmanager:$(ALERTMANAGER_VERSION) -f Dockerfile .     

test:
	go test ./... -v

vet:
	go vet ./... 
