# Run In Parallel

## What it does
Rip will run any program you like in parallel. It starts as many processes as you like and lets you define the arguments for the program.

## Build
To build, you need a running copy of [Go](http://golang.org). Then just buid by using:
```
go build rip.go
```

## Usage
Usage e.g.: `./rip -n 32 -c "date" -- +%S`
- `-n 32` number of processes to start, e.g. 32
- `-c date` process to start, e.g. date
- `-- +%S` arguments given to program have to be behind `--`.
