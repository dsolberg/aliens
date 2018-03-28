run:
	go run main.go -aliens=10000

build:
	go clean
	go build -o dist/simulator
	chmod +x dist/simulator


