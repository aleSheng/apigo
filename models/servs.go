package models

import (
	"github.com/anlint/apigo/models/mymongo"
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Serv struct {
	ID       bson.ObjectId    `bson:"_id"      json:"id,omitempty"`
	Title     string    `bson:"title"     json:"title,omitempty"`
	Create_at time.Time `bson:"create_at" json:"create_at"`
	Pic       string    `bson:"pic" json:"pic,omitempty"`
}
type Servlist struct {
	Servs []Serv
}

func Findservbyid(id string) (u Serv, err error) {
	mConn := mymongo.Conn()
	defer mConn.Close()
	c := mConn.DB("anlintdb").C("servads")
	err = c.Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&u)
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
