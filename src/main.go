package main

import "github.com/gin-gonic/gin"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
//import "svc"

type Whitelist struct {
     IP []byte `"json:ip" "bson:ip"`
}

func main() {
     session, err := mgo.Dial("localhost:27017")
     if err != nil {
          panic(err)
     }
     defer session.Close()
     db := session.DB("avoxi_challenge")

     r := gin.Default()
     r.GET("/login", func(c *gin.Context) {
          clientIP := c.ClientIP()
          whitelist := []Whitelist{}
          db.C("Whitelist").Find(bson.M{}).All(&whitelist)
          /*
          if ( ! svc.isAuthorized( clientIP, whitelist ) ) {
               c.JSON(403, gin.H{
                    "login failure": "client ip address is unauthorized",
               })
          }
          */
          if whitelist != nil && clientIP != "" {
               c.JSON(200, gin.H{
                    "clientIP": clientIP,
                    "whitelist": whitelist,
               })
          }
     })
     r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}