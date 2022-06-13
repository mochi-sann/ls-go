package main

import (
	"fmt"
	"log"
	"os"
)

func readCurrentDir() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}

	defer file.Close()

	fileList, _ := file.Readdir(0)

	for _, files := range fileList {
		fmt.Printf("\n%v ", files.Name())
	}

}
func main() {
	readCurrentDir()
}
