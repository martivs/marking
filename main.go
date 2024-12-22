package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type RW struct {
	fileR   *os.File
	fileW   *os.File
	scanner *bufio.Scanner
	writer  *bufio.Writer
	line    string
	time    string
}

func main() {

	var rw RW

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
	rw.fileR = fileR
	defer fileR.Close()

	// create fileW
	fileW, err := os.Create(currentTime + ".dxf")
	if err != nil {
		log.Fatal("Error creating file:", err)
		return
	}
	rw.fileW = fileW
	defer fileW.Close()

	// create scanner & writer
	rw.scanner = bufio.NewScanner(fileR)
	rw.writer = bufio.NewWriter(fileW)

	// work
	work(&rw, currentTime)
}

func work(rw *RW, currentTime string) {

	// read line by line and write to file
	for rw.scanner.Scan() {
		line := rw.scanner.Text()

		_, err := rw.writer.WriteString(line + "\n")
		if err != nil {
			os.Remove(currentTime + ".dxf")
			log.Fatal("Error writing to file:", err)
		}
	}

	// check for errors
	if err := rw.scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	/*
		Убедиться, что все данные из буфера записаны в файл
		+ обработка ошибок для writer.Flush()
	*/
	if err := rw.writer.Flush(); err != nil {
		log.Fatal("Error flushing writer:", err)
	}
}
