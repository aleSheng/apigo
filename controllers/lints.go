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
func (this *LintController) Devote() {
//	id:=this.GetString(":id")
//	obs,err := models.Findlintbyid(id)
//	if(err!=nil){
//		this.Data["json"] = err
//	}else{
//		this.Data["json"] = "{'err':'0','msg':'" + obs.ID+"'}"
//	}
	this.Data["json"] = map[string]string{
		"error": "0",
		"msg":  "举报已提交",
	}
	this.ServeJson()
}

func (this *LintController) Getall() {
	lastdate, er:= time.Parse("2006-01-02T15:04:05Z",this.GetString("lastdate"))
	if er != nil {
		lastdate = time.Now()
	}
	cateid, err := this.GetInt("cateid")
	if err != nil {
		cateid=0
	}
	obs := models.Getlints(lastdate,cateid)
	this.Data["json"] = obs
	this.ServeJson()
}

