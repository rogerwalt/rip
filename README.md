# Run In Parallel

## What it does
Rip will run any program you like in parallel. It starts as many processes as you like and lets you define the arguments for the program as well as a different seed for every process. Especially useful for heuristic simulations like monte carlo and such.

## Build
To build, you need a running copy of [Go](http://golang.org). Then just build by using:
```
go build rip.go
```

## Usage
Usage e.g.: `./rip -n 32 -c echo -- _seed_`:
```
> ./rip -n 3 -c echo -- hello from _seed_
hello from 0
hello from 2
hello from 1
```
- `-n 3` number of processes to start
- `-c date` program to start
- `-- _seed_` arguments given to program have to be behind `--`, `_seed_` is replaced by the process number
