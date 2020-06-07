package api

import "github.com/gin-gonic/gin"

func loginAPI(c *gin.Context) {
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
}