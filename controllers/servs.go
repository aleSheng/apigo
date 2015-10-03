package controllers

import (
	"github.com/anlint/apigo/models"
	"time"
)

type ServController struct {
	BaseController
}
func (this *ServController) Getone() {
	id:=this.GetString(":id")
	obs , err := models.Findservbyid(id)
	if(err!=nil){
		this.Data["json"] = err
	}else{
		this.Data["json"] = obs
	}
	this.ServeJson()
}

func (this *ServController) Getall() {
	lastdate, error:= time.Parse("2006-01-02T15:04:05Z",this.GetString("lastdate"))
	if error != nil {
		lastdate = time.Now()
	}
	obs := models.Getallserv(lastdate)
	this.Data["json"] = obs
	this.ServeJson()
}

