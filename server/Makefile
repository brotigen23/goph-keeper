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


.PHONY: test coverage