package mongo

import "gopkg.in/mgo.v2"

func createConnection() *mgo.Session {
	session, err := mgo.Dial("localhost:27017")
     if err != nil {
          panic(err)
     }
     return session
}

