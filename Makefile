run:
	go run main.go

build:
	go clean
	go build -o dist/simulator
	chmod +x dist/simulator


