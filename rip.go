package main

import (
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func main() {
	// initialize commandline flags
	numProc := flag.Int("n", 4, "number of processes")
	wordCmd := flag.String("c", "", "command")

	// parse flags
	flag.Parse()

	// check for empty command
	if (*wordCmd == "") {
		log.Fatalf("Error: Command not given")
		os.Exit(2)
	}

	args := flag.Args()

	// check for _seed_ in tail save index position of seed
	seedindex := -1

	for index, element := range args {
		if element == "_seed_" {
			seedindex = index
		}
	}

	// launch as many processes as wanted
	for i := 0; i < *numProc; i++ {
		// replace _seed_ by i
		if seedindex != -1 {
			args[seedindex] = strconv.Itoa(i)
		}

		// the tail is given as arguments to the command which is run
		cmd := exec.Command(*wordCmd, args...)

		// create stdout, stderr streams of type io.Reader
		stdout, err := cmd.StdoutPipe()
		checkError(err)
		stderr, err := cmd.StderrPipe()
		checkError(err)

		// non-blockingly echo command output to terminal
		go io.Copy(os.Stdout, stdout)
		go io.Copy(os.Stderr, stderr)

		// start command
		err = cmd.Start()
		checkError(err)

		// don't let main() exit before our command has finished running
		defer cmd.Wait()  // doesn't block
	}
}