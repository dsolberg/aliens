DEFAULT_ALIENS = 2000

run:
	go run main.go -aliens=$(DEFAULT_ALIENS)

build:
	go clean
	go build -o dist/simulator
	chmod +x dist/simulator

test:
	tests/pacifistaliens.sh $(DEFAULT_ALIENS)


