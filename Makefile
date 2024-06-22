clean:
	rm alertmanager

build:
	go build

test:
	go test ./... -v