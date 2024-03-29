package models

import (
	"github.com/astaxie/beego"
	"github.com/anlint/apigo/models/mymongo"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Lint struct {
	ID       bson.ObjectId    `bson:"_id"      json:"id,omitempty"`
	Text     string    `bson:"text"     json:"text,omitempty"`
	Create_at time.Time `bson:"create_at" json:"create_at"`
	Pic       string    `bson:"pic" json:"pic,omitempty"`
	Cateid    int     `bson:"cateid" json:"cateid"`
	Styleid   int     `bson:"styleid" json:"styleid"`
	Pubname   string  `bson:"pubname" json:"pubname"`
	Headimg   string  `bson:"headimg" json:"headimg"`
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

	dbname :=beego.AppConfig.String("mongodb::dbname")
	c := mConn.DB(dbname).C("lints")
	iter := c.Find(bson.M{"create_at":bson.M{"$lt":lastdate}}).Sort("-create_at").Skip(0).Limit(24).Iter()
	var result Lint
	for iter.Next(&result) {
		personAll.Lints = append(personAll.Lints, result)
	}
	return
}