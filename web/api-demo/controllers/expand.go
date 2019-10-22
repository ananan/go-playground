package controllers

import "github.com/astaxie/beego"

type ExpandController struct {
	beego.Controller
}

func (self *ExpandController) Get() {
	var result ShortResult
	shorturl := self.Input().Get("shorturl")
	result.UrlShort = shorturl

	if urlcache.IsExist(shorturl) {
		result.UrlLong = urlcache.Get(shorturl).(string)
	} else {
		result.UrlLong = ""
	}
	self.Data["json"] = result
	self.ServeJSON()
}
