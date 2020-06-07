package api

import "github.com/gin-gonic/gin"
import "gopkg.in/mgo.v2"

func InitRoutes(r *gin.Engine, db *mgo.Database) {
	
	r.GET("/login", func(c *gin.Context) {
		loginAPI(c, db)
	})
}