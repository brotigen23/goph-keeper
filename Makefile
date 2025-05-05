BUILDDIR=bin

all:
	cd client/
	make

build:
	go build -o ${BUILDDIR}/client cmd/client/main.go

create-docker-images:


docker-image-client:
	docker build --tag goph-keeper-client ./client


test:
	cd client/ && go test ./...
	cd server && go test ./...

.PHONY: all build all docker-image-client

swag:
	~/go/bin/swag init \
    -g ./server/cmd/server/main.go \
	--output ./docs \
    --parseDependency \
    --parseInternal \
    --parseDepth 3