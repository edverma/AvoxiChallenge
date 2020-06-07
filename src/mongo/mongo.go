package mongo

import (
     "log"

     "gopkg.in/mgo.v2"
     "gopkg.in/mgo.v2/bson"
     "github.com/edverma/AvoxiChallenge/models"
)

func CreateConnection() (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.Dial("localhost:27017")
     db := session.DB("avoxi_challenge")
     return session, db, err
}

func InitData(db *mgo.Database) {
     sample := models.Whitelist{}
     err := db.C("Whitelist").Find(bson.M{}).One(&sample)
     if err != nil {
          log.Print(err)
     }
     if sample.IsoCode == "" {
          err := db.C("Whitelist").Insert(
               bson.M{"iso_code": "US"},
               bson.M{"iso_code": "IQ"},
               bson.M{"iso_code": "WS"},
               bson.M{"iso_code": "BW"},
               bson.M{"iso_code": "AX"},
          )
          if err != nil {
               log.Fatal(err)
          }
     }
}

