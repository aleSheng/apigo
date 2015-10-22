package models

import (
	"github.com/anlint/apigo/models/mymongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Serv struct {
	ID       bson.ObjectId    `bson:"_id"      json:"id,omitempty"`
	Title     string    `bson:"title"     json:"title,omitempty"`
	Create_at time.Time `bson:"create_at" json:"create_at"`
	Pic       string    `bson:"pic" json:"pic,omitempty"`
	Link       string    `bson:"link" json:"link,omitempty"`
}
type Servlist struct {
	Servs []Serv
}

func Findservbyid(id string) (u Serv, err error) {
	mConn := mymongo.Conn()
	defer mConn.Close()
	c := mConn.DB("anlintdb1").C("servads")
	err = c.Find(bson.M{"_id":bson.ObjectIdHex(id)}).One(&u)
	return
}

func  Getallserv(lastdate time.Time) (personAll Servlist) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb1").C("servads")
	iter := c.Find(bson.M{"create_at":bson.M{"$lt":lastdate}}).Skip(0).Limit(24).Iter()
	var result Serv
	for iter.Next(&result) {
		personAll.Servs = append(personAll.Servs, result)
	}
	return
}
