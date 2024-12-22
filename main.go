package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// get current time
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	currentTime = strings.ReplaceAll(currentTime, "-", "")
	currentTime = strings.ReplaceAll(currentTime, ":", "")
	currentTime = strings.ReplaceAll(currentTime, " ", "")
	currentTime = currentTime[2 : len(currentTime)-2]

	// open fileR
	fileR, err := os.Open("patt.dxf")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer fileR.Close()

	// create fileW
	fileW, err := os.Create(currentTime + ".dxf")
	if err != nil {
		log.Fatal("Error creating file:", err)
		return
	}
	defer fileW.Close()

	// make a scanner
	scanner := bufio.NewScanner(fileR)

	// create a writer
	writer := bufio.NewWriter(fileW)

	// read line by line
	for scanner.Scan() {
		line := scanner.Text()

		_, err := writer.WriteString(line + "\n")
		if err != nil {
			os.Remove(currentTime + ".dxf")
			log.Fatal("Error writing to file:", err)
		}

	}

	// check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
