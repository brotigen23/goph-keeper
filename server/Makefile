all:
	go run cmd/server/main.go

clean:
	rm log/*

test:
	go test ./... -v -cover

testCover:
	go test ./... -coverprofile test/cover/coverage.out -count=1
	go tool cover -html=test/cover/coverage.out -o test/cover/cover.html

coverage:
	@go test -coverprofile=coverage.out ./... -count=1
	@echo "Checking total coverage..."
	@if [ -f coverage.out ]; then \
		echo "Total coverage: $$(go tool cover -func=coverage.out | grep 'total' | awk '{print $$3}')"; \
	else \
		echo "coverage.out file not found"; \
	fi

mockgen:
	~/go/bin/mockgen -source=./internal/repository/repository.go -destination=./internal/repository/mock/userRepository.go -package=mock UserRepo


swag:
	rm -rf ./docs/
	~/go/bin/swag init \
    -g ./cmd/server/main.go \
	--output ./docs \
    --parseDependency \
    --parseInternal \
    --parseDepth 3

.PHONY: test coverage mockgen
