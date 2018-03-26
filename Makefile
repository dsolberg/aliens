.PHONY: build
build:
	go clean
	go build -o dist/simulator
	chmod +x dist/simulator


