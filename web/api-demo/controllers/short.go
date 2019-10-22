package controllers

import (
	"api-demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
)

var urlcache cache.Cache

func init() {
	urlcache, _ = cache.NewCache("memory",
		`{"interval": 0`)
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

type ShortController struct {
	beego.Controller
}

func (self *ShortController) Get() {
	var result ShortResult
	longurl := self.Input().Get("longurl")
	result.UrlLong = longurl

	urlmd5 := models.GetMD5(longurl)
	logs.Info(urlmd5)

	if urlcache.IsExist(urlmd5) {
		result.UrlShort = urlcache.Get(urlmd5).(string)
	} else {
		result.UrlShort = models.Generate()
		if err := urlcache.Put(urlmd5, result.UrlShort, 0); err != nil {
			logs.Warn(err)
		}
		if err := urlcache.Put(result.UrlShort, longurl, 0); err != nil {
			logs.Warn(err)
		}
	}
	logs.Warn("url-long: ", result.UrlLong, "url-short: ", result.UrlShort)
	self.Data["json"] = result
	self.ServeJSON()
}
