all:
	go run cmd/server/main.go

clean:
	rm log/*

test:
	go test ./... -v -cover

testCover:
	go test ./... -coverprofile test/cover/coverage.out -count=1
	go tool cover -html=test/cover/coverage.out -o test/cover/cover.html


.PHONY: test