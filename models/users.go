package models

import (
	"github.com/anlint/apigo/models/mymongo"
	"fmt"
)

type User struct {
	ID       string    `bson:"_id"      json:"_id,omitempty"`
	Name     string    `bson:"name"     json:"name,omitempty"`
}
const PW_HASH_BYTES = 64
type Men struct {
	Users []User
}

func FinduserById() (u User) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")
	c.Find(nil).One(u)
	return
}

func  Getallusers() (personAll Men) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("anlintdb").C("users")
	iter := c.Find(nil).Iter()
	var result User
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result.Name)
		personAll.Users = append(personAll.Users, result)
	}
	return
}
