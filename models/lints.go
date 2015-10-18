package models

import (
	"github.com/anlint/apigo/models/mymongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Lint struct {
	ID       bson.ObjectId    `bson:"_id"      json:"id,omitempty"`
	Title     string    `bson:"title"     json:"title,omitempty"`
	Create_at time.Time `bson:"create_at" json:"create_at"`
	Pic       string    `bson:"pic" json:"pic,omitempty"`
	Cateid    int     `bson:"cateid" json:"cateid"`
	Styleid   int     `bson:"styleid" json:"styleid"`
}
type Lintlist struct {
	Lints []Lint
}

func Findlintbyid(id string) (u Lint, err error) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb1").C("lints")
	err = c.FindId(bson.ObjectIdHex(id)).One(&u)
	return
}

func  Getlints(lastdate time.Time, cateid int) (personAll Lintlist) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb1").C("lints")
	iter := c.Find(bson.M{"create_at":bson.M{"$lt":lastdate}}).Skip(0).Limit(12).Iter()
	var result Lint
	for iter.Next(&result) {
		personAll.Lints = append(personAll.Lints, result)
	}
	return
}