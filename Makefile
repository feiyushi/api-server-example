build:
	go build -o bin/apiserver apiserver/cmd

test:
	go test ./...

clean:
	rm -rf bin