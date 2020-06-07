package main

import (
     "log"

     "github.com/gin-gonic/gin"
     "github.com/edverma/AvoxiChallenge/mongo"
     "github.com/edverma/AvoxiChallenge/api"
)

func main() {
     session, db, err := mongo.CreateConnection()
     if err != nil {
          log.Fatal(err)
          return
     }
     defer session.Close()

     mongo.InitData(db)

     r := gin.Default()
     api.InitRoutes(r, db)
     r.Run() //localhost:8080
}