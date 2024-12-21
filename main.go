package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	currentTime = strings.ReplaceAll(currentTime, "-", "")
	currentTime = strings.ReplaceAll(currentTime, ":", "")
	currentTime = strings.ReplaceAll(currentTime, " ", "")
	currentTime = currentTime[2 : len(currentTime)-2]
	fmt.Println(currentTime)

}
