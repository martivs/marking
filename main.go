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

	// open file
	file, err := os.Open("patt.dxf")
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer file.Close()

	// make a scanner
	scanner := bufio.NewScanner(file)

	// read line by line
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "TEXT" {
			fmt.Println(i, "TEXT")
			i++
		}
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
