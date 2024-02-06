package logging

import (
	"fmt"
	"time"
)
const (
	InfoStart string = "|INFO|"
	ErrorStart string = "|WARN|"
	SuccessStart string = "|SUCC|"
	FailStart string = "|FAIL|"
)

func InfoMessage(funcName string){
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s - %s is called\n",InfoStart, timestamp, funcName)
}

func ErrorMessage(funcName string, errorMsg string){
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s - %s is throw error\n",ErrorStart, timestamp, funcName)
	fmt.Printf("%s Error Message is:\n",ErrorStart)
	fmt.Printf("%s %s\n", ErrorStart, errorMsg)
}

func SuccessMessage(funcName string){
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s - %s completed successfully\n", SuccessStart, timestamp, funcName)
}

func FailMessage(funcName string, reason string){
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Printf("%s %s - %s completed unsuccessfully\n",FailStart, timestamp, funcName)
	fmt.Printf("%s %s - Reason is: %s",FailStart, timestamp, reason)
}