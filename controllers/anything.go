package controllers

import(
	"beegoapi/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AnythingController struct{
	beego.Controller
}


func (a *AnythingController) Post() {
	logs.Info("HTTP Post AddAnything")
	var anything models.Anything
	json.Unmarshal(a.Ctx.Input.RequestBody, &anything)
	uid := models.AddAnything(anything)
	a.Data["json"] = map[string]string{"uid": uid}
	a.ServeJSON()
}

// @Title GetAll
// @Description get all anythings
// @Success 200 {anything} models.Anything
// @Failure 403 :anythingId is empty
// @router / [get]
func (a *AnythingController) GetAll() {
	logs.Info("HTTP Get GetAnythingMap")
	anythings := models.GetAnythingMap()
	a.Data["json"] = anythings
	a.ServeJSON()
}

// @Title Get
// @Description find anything by anythingId
// @Param	anythingId		path 	string	true		"the anythingId you want to get"
// @Success 200 {anything} models.Anything
// @Failure 403 :anythingId is empty
// @router /:anythingId [get]
func (a *AnythingController) Get() {
	anythingId := a.Ctx.Input.Param(":anythingId")
	logs.Info("HTTP Get GetAnything - anythingId : %s",anythingId)	
	if anythingId != "" {
		at, err := models.GetAnything(anythingId)
		if err != nil {
			a.Data["json"] = err.Error()
		} else {
			a.Data["json"] = at
		}
	}
	a.ServeJSON()
}