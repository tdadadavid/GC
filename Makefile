build:
	go build -o gc

run:
	go run main.go

test:
	go test -v ./...

clean:
	rm -f gc
