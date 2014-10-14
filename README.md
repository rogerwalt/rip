# Run In Parallel

## What it does
Rip will run any program you like in parallel. It starts as many processes as you like and lets you define the arguments for the program.

## Build
To build, you need a running copy of [Go](http://golang.org). Then just buid by using:
```go build rip.go
```

## Usage

Usage e.g.: `./rip -n 32 -c "date" -- +%S`