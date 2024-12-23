package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type RW struct {
	fileR       *os.File
	fileW       *os.File
	scanner     *bufio.Scanner
	writer      *bufio.Writer
	currentTime string
	name        string
}

func main() {

	var rw RW
	rw.name = "SBS07H"

	// get current time
	rw.currentTime = time.Now().Format("2006-01-02 15:04:05")
	rw.currentTime = strings.ReplaceAll(rw.currentTime, "-", "")
	rw.currentTime = strings.ReplaceAll(rw.currentTime, ":", "")
	rw.currentTime = strings.ReplaceAll(rw.currentTime, " ", "")
	rw.currentTime = rw.currentTime[2 : len(rw.currentTime)-2]

	// open fileR
	fileR, err := os.Open("patt.dxf")
	if err != nil {
		log.Fatal(err)
		return
	}
	rw.fileR = fileR
	defer fileR.Close()

	// create fileW
	fileW, err := os.Create(rw.name + "_" + rw.currentTime + ".dxf")
	if err != nil {
		log.Fatal(err)
		return
	}
	rw.fileW = fileW
	defer fileW.Close()

	// create scanner & writer
	rw.scanner = bufio.NewScanner(fileR)
	rw.writer = bufio.NewWriter(fileW)

	// work
	work(&rw)
}

func work(rw *RW) {

	// read line by line and write to file
	for rw.scanner.Scan() {
		line := rw.scanner.Text()
		if line != "TEXT" {
			writeLine(rw, line)
		} else {
			lines := readTEXT(rw)
			writeTEXT(rw, lines, "TIME")
		}
	}

	// check for errors
	if err := rw.scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func writeLine(rw *RW, line string) {

	_, err := rw.writer.WriteString(line + "\n")
	if err != nil {
		log.Fatal(err)
	}

	// Убедиться, что все данные из буфера записаны в файл
	// + обработка ошибок для writer.Flush()
	if err := rw.writer.Flush(); err != nil {
		log.Fatal(err)
	}

}

func readTEXT(rw *RW) []string {

	var lines []string
	lines = append(lines, "TEXT")

	for rw.scanner.Scan() {
		if rw.scanner.Text() == "  0" {
			lines = append(lines, "  0")
			break
		} else {
			lines = append(lines, rw.scanner.Text())
		}
	}

	// check for errors
	if err := rw.scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func writeTEXT(rw *RW, lines []string, flag string) {

	for i, line := range lines {
		switch {
		case i > 0 && lines[i-1] == "  1":
			if flag == "NAME" {
				writeLine(rw, rw.name)
			}
			if flag == "TIME" {
				writeLine(rw, rw.currentTime)
			}
		case i > 0 && lines[i-1] == " 21":
			writeLine(rw, line)
		case i > 0 && lines[i-1] == " 40":
			writeLine(rw, "3.0")
		default:
			writeLine(rw, line)
		}
	}

}
