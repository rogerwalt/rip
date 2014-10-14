package main

import (
	"fmt"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
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

	// print flag variables
	fmt.Println("n:", *numProc)
	fmt.Println("c:", *wordCmd)
	fmt.Println("tail:", flag.Args())

	// check for empty command
	if (*wordCmd == "") {
		log.Fatalf("Error: Command not given")
		os.Exit(2)
	}

	// launch as many processes as wanted
	for i := 0; i < *numProc; i++ {
		// the tail is given as arguments to the command which is run
	    cmd := exec.Command(*wordCmd, flag.Args()...)

	    // create stdout, stderr streams of type io.Reader
	    stdout, err := cmd.StdoutPipe()
	    checkError(err)
	    stderr, err := cmd.StderrPipe()
	    checkError(err)

	    // start command
	    err = cmd.Start()
	    checkError(err)

	    // don't let main() exit before our command has finished running
	    defer cmd.Wait()  // doesn't block

	    // non-blockingly echo command output to terminal
	    go io.Copy(os.Stdout, stdout)
	    go io.Copy(os.Stderr, stderr)
	}
}