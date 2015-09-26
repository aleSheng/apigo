package models

import (
	"github.com/anlint/apigo/models/mymongo"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Serv struct {
	ID       bson.ObjectId    `bson:"_id"      json:"id,omitempty"`
	Title     string    `bson:"title"     json:"title,omitempty"`
}
type Servlist struct {
	Servs []Serv
}

func Findservbyid() (u Serv) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb").C("servads")
	c.Find(nil).One(u)
	return
}

func  Getallserv() (personAll Servlist) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb").C("servads")
	iter := c.Find(nil).Iter()
	var result Serv
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result.ID)
		personAll.Servs = append(personAll.Servs, result)
	}
	return
}
