all: clean
	go build -o bin/term main.go

clean:
	rm -rf bin/*