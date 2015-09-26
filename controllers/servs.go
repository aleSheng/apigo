package controllers

import (
	"github.com/anlint/apigo/models"
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
	obs := models.Getallserv()
	this.Data["json"] = obs
	this.ServeJson()
}

