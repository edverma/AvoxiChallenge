package services

import (
	"log"

	geolite2 "github.com/oschwald/geoip2-golang"
	"github.com/edverma/AvoxiChallenge/models"
)

type Session struct {
	Authorized bool
}

func NewSessionService() *Session{
	return &Session{false}
}

func (session *Session) IsAuthorized (ip []byte, whitelist []models.Whitelist) bool {
	geoLiteDB, err := geolite2.Open("../GeoLite2-Country/GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer geoLiteDB.Close()

	record, err := geoLiteDB.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	isoCode := record.Country.IsoCode

	for _, item := range whitelist {
		if item.IsoCode == isoCode {
			session.Authorized = true
			return true
		}
	}
	session.Authorized = false
	return false
}