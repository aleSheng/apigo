package controllers

import (
	"github.com/anlint/apigo/models"
)

type LintController struct {
	BaseController
}
func (this *LintController) Getone() {
	id:=this.GetString(":id")
	obs,err := models.Findlintbyid(id)
	if(err!=nil){
		this.Data["json"] = err
	}else{
		this.Data["json"] = obs
	}
	this.ServeJson()
}


func (this *LintController) Getall() {
	obs := models.Getalllint()
	this.Data["json"] = obs
	this.ServeJson()
}

