package main

import (
	"github.com/gin-gonic/gin"

	"typicalypetprojects/pkg/shorterurl"
	"typicalypetprojects/pkg/typing"
	"typicalypetprojects/pkg/database"
	
)


func main() {

	var existUrls = []typing.Urls{}

    db, err := database.ConnectPg()
    if err != nil{
        panic(err)
    }
    defer db.Close()

	router:= gin.Default()
	router.POST(
		"/shortMyUrl", 
		func(c *gin.Context){
			shorterurl.PostUrl(c, &existUrls)
		},
	)
	
	router.Run("localhost:8080")
}
