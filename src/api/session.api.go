package api

import ( 
	"net"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/edverma/AvoxiChallenge/models"
	"github.com/edverma/AvoxiChallenge/services"
)

func loginAPI(c *gin.Context, db *mgo.Database) {
	sessionSvc := services.NewSessionService()
	//clientIP := net.ParseIP( c.ClientIP() )
	clientIP  := net.ParseIP( "71.204.91.63" )
	whitelist := []models.Whitelist{}

	err := db.C("Whitelist").Find(nil).All(&whitelist)
	if err != nil {
		log.Fatal(err)
	}
	if ( ! sessionSvc.IsAuthorized( clientIP, whitelist ) ) {
		c.JSON(403, gin.H{
			"login failure": "client ip address is unauthorized",
		})
	} else {
		c.JSON(200, "authorized")
	}
}