package main

import (
    "fmt"
    "github.com/gin-gonic/gin" 
)

func main() {
    fmt.Println("🚀 iBeaconRoll server started!")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello from iBeaconRoll 🎉",
        })
    })
	
	r.Run(":8080") 

}