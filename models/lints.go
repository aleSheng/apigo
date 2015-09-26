package models

import (
	"github.com/anlint/apigo/models/mymongo"
	"fmt"
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

func Findlintbyid() (u Lint) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb").C("lints")
	c.Find(nil).One(u)
	return
}

func  Getalllint() (personAll Lintlist) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb").C("lints")
	iter := c.Find(nil).Iter()
	var result Lint
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result.ID)
		personAll.Lints = append(personAll.Lints, result)
	}
	return
}
