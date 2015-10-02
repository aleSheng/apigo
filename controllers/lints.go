package controllers

import (
	"github.com/anlint/apigo/models"
	"github.com/astaxie/beego"
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
	lastdate, error:= time.Parse("20060102 15:04:05",this.GetString("lastdate"))
	if error != nil {
		lastdate = time.Now()
	}
	cateid, error := this.GetInt("cateid")
	if error != nil {
		cateid=0
	}
	beego.Debug(lastdate)
	beego.Debug(cateid)
	obs := models.Getlints(lastdate,cateid)
	this.Data["json"] = obs
	this.ServeJson()
}

