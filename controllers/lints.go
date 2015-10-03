package controllers

import (
	"github.com/anlint/apigo/models"
	"time"
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
	lastdate, error:= time.Parse("2006-01-02T15:04:05Z",this.GetString("lastdate"))
	if error != nil {
		lastdate = time.Now()
	}
	cateid, error := this.GetInt("cateid")
	if error != nil {
		cateid=0
	}
	obs := models.Getlints(lastdate,cateid)
	this.Data["json"] = obs
	this.ServeJson()
}

