package controllers

import (
	"blog/app/models"
	"blog/app/routes"
	"blog/app/support"
	"encoding/json"
	"github.com/revel/revel"
	"strconv"
)

type Login struct {
	*revel.Controller
}

func (l Login) SignIn() revel.Result {
	return l.Render()
}

func (l Login) SignInHandler(name, passwd string) revel.Result {

	model := &models.Admin{Name: name, Passwd: passwd}
	admin, err := model.SignIn(l.Request)

	if err != "" {
		revel.ERROR.Println(err)
		l.Flash.Error("msg", err)
		return l.Redirect(routes.Login.SignIn())
	}

	revel.INFO.Println(admin)

	l.Session["UID"] = strconv.Itoa(admin.Id)

	data, _ := json.Marshal(&admin)
	support.Cache.Set(support.SPY_ADMIN_INFO+strconv.Itoa(admin.Id), string(data), 0)

	return l.RenderHtml(l.Session.Id())
}
