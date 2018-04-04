# Alien Invasion

## Minimum requirements

Requirement|Notes
---|---
Go version | Go1.9 or higher

## Install

To install from source:

```
go get -u github.com/dsolberg/aliens
```

To get to the local code:
```
cd $GOPATH/src/github.com/dsolberg/aliens
git checkout tags/1.5.1
```

To execute with the default parameters of 2000 aliens run:

```
make
```

To execute with non-standard parameters run:

```
go run main.go -aliens=<number of aliens>
```

To run tests run:
```
make test
```

To compile and run execute:

```
make build
dist/simulator
```
