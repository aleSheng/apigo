package controllers

import (
	"github.com/anlint/apigo/models"
)

type LintController struct {
	BaseController
}
func (this *LintController) Getone() {
	obs := models.Findlintbyid()
	this.Data["json"] = obs
	this.ServeJson()
}


func (this *LintController) Getall() {
	obs := models.Getalllint()
	this.Data["json"] = obs
	this.ServeJson()
}

