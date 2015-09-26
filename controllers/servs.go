package controllers

import (
	"github.com/anlint/apigo/models"
)

type ServController struct {
	BaseController
}
func (this *ServController) Getone() {
	obs := models.Findservbyid()
	this.Data["json"] = obs
	this.ServeJson()
}


func (this *ServController) Getall() {
	obs := models.Getallserv()
	this.Data["json"] = obs
	this.ServeJson()
}

