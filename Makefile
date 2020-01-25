build:
	go build -o qdc cmd/qdc/main.go

run: build
	./qdc

install: build
	cp qdc /usr/local/bin

uninstall:
	rm /usr/local/bin/qdc

clean:
	rm qdc
