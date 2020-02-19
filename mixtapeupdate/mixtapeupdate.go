package main

import (
	"flag"
	"log"
	"os"

	"../pkg/mixtape"
)

var stdLog = log.New(os.Stdout, "main: ", 0)

func printUsage() {
	stdLog.Println("params:")
	stdLog.Println("-mixTape")
	stdLog.Println("-changes")
	stdLog.Println("-output (Optional)")
}

var (
	mixTapeFilePath = flag.String("mixTape", "", "specify path to the mix tape json file")
	changesFilePath = flag.String("changes", "", "specify path to the changes json file")
	outputFilePath  = flag.String("output", "", "specify path to the output json file")
)

func main() {
	flag.Parse()
	defaultOutputPath := "output.json"

	if (*mixTapeFilePath == "") || (*changesFilePath == "") {
		printUsage()
		return
	}
	if *outputFilePath == "" {
		outputFilePath = &defaultOutputPath
	}

	err := mixtape.HandleChangesToMixTape(*mixTapeFilePath, *changesFilePath, *outputFilePath)
	if err != nil {
		stdLog.Printf("The following error occurred when applying changes to mixtape: %v", err)
	} else {
		stdLog.Printf("Successfully wrote output file: %s", *outputFilePath)
	}
}
