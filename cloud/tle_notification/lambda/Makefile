build: cleanup
	env GOOS=linux GOARCH=amd64 go build -o ./build

cleanup:
	rm -rf ./build
	mkdir ./build

test:
	go test ./...
