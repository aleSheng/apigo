package controllers

import (
	"github.com/anlint/apigo/models"
)

type UserController struct {
	BaseController
}
func (this *UserController) Getone() {
	obs := models.FinduserById()
	this.Data["json"] = obs
	this.ServeJson()
}


func (this *UserController) Getall() {
	obs := models.Getallusers()
	this.Data["json"] = obs
	this.ServeJson()
}

