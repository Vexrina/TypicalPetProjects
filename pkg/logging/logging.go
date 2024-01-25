package logging

import (
	"fmt"
	"time"
)
const (
	InfoStart string = "|INFO|"
	// InfoStart string = "|INFO|"
)
func InfoMessage(funcName string){
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s - %s is called\n",InfoStart, timestamp, funcName)
}