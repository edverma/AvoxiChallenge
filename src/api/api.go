package api

import "github.com/gin-gonic/gin"
import "gopkg.in/mgo.v2"
import "mongo"

func Init(db *mgo.Database) {
	r := gin.Default()
	
	r.GET("/login", func(c *gin.Context) {
		loginAPI(c, db)
	})

	r.Run() //localhost:8080
}