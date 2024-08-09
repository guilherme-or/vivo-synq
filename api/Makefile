go-test:
	go test -v ./test/...

go-build:
	go build -o ./bin/api ./cmd/server.go

run: go-build
	./bin/api
