package models

import (
	"github.com/astaxie/beego"
	"github.com/anlint/apigo/models/mymongo"
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId    `bson:"_id"      json:"_id,omitempty"`
	Name     string    `bson:"name"     json:"name,omitempty"`
	Create_at time.Time `bson:"create_at"`
}
const PW_HASH_BYTES = 64
type Men struct {
	Users []User
}

func FinduserById() (u User) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb1").C("users")
	c.Find(nil).One(u)
	return
}

func  Getallusers() (personAll Men) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	dbname :=beego.AppConfig.String("mongodb::dbname")
	c := mConn.DB(dbname).C("users")
	iter := c.Find(nil).Iter()
	var result User
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result.ID)
		personAll.Users = append(personAll.Users, result)
	}
	return
}
