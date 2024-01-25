package main

import (
	"fmt"
	"time"
	"typicalpetprojects/pkg/logging"
	// "typicalypetprojects/pkg/shorterurl"
	// "net/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	timestamp := time.Now().UTC()

	fmt.Println(timestamp.Format("2006-01-02 15:04:05"))

	logging.InfoMessage("someFunction")
}
