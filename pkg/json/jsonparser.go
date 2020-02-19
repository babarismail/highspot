package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"../types"
)

var stdLog = log.New(os.Stdout, "json parser: ", 0)

func readBytes(filePath string) ([]byte, error) {
	if !fileExists(filePath) {
		return nil, fmt.Errorf("file path : %s doesn't exist", filePath)
	}
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		stdLog.Printf("Could not read file: %s", filePath)
		return nil, err
	}
	return fileBytes, nil
}

func fileExists(filename string) (exists bool) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//ReadMixTape reads the file at the given file path and returns the MixTape struct
func ReadMixTape(filePath string) (mixTape types.MixTape, err error) {
	mixTapeBytes, err := readBytes(filePath)
	if err != nil {
		return mixTape, err
	}
	err = json.Unmarshal(mixTapeBytes, &mixTape)
	if err != nil {
		stdLog.Printf("error in JSON unmarshal: %s", filePath)
		return mixTape, err
	}
	return mixTape, nil
}

//ReadChanges reads the file at the given file path and returns the Changes struct
func ReadChanges(filePath string) (changes types.Changes, err error) {
	changesBytes, err := readBytes(filePath)
	if err != nil {
		return changes, err
	}
	err = json.Unmarshal(changesBytes, &changes)
	if err != nil {
		stdLog.Printf("error in JSON unmarshal: %s", filePath)
		return changes, err
	}
	return changes, nil
}

//WriteOutputToFile writes the mix tape file to the specified file path
func WriteOutputToFile(outputFilePath string, mixTape types.MixTape) error {
	//writing with indentation for better readability
	mixTapeBytes, err := json.MarshalIndent(mixTape, "", "  ")
	if err != nil {
		return err
	}
	//not too fixated on file permissions. this will depend on requirements on what the correct file permissions should be
	err = ioutil.WriteFile(outputFilePath, mixTapeBytes, 0644)
	return err
}
