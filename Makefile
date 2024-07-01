ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ALERTMANAGER_VERSION := $(shell ./alertmanager --version | cut -d" " -f3)

sed:
	sed -i 's/WEBHOOK_SECRET/${WEBHOOK_SECRET}/' alert-manager-config.yml
	sed -i 's/WEBHOOK_SECRET/${WEBHOOK_SECRET}/' deployment/toy_alert_manager.yml
	# sed 's/WEBHOOK_SECRET/${WEBHOOK_SECRET}/' alert-manager-config.yml

clean: 
	rm alertmanager

build: vet
	go build -tags netgo

docker-build: build
	 docker build -t alertmanager:$(ALERTMANAGER_VERSION) -f Dockerfile .

docker-push: docker-build
	docker tag alertmanager:$(ALERTMANAGER_VERSION) sohom83/tam:$(ALERTMANAGER_VERSION)
	docker push sohom83/tam:$(ALERTMANAGER_VERSION)

test:
	go test ./... -v

vet:
	go vet ./... 
