package controllers

import (
	"encoding/json"
	"little_project/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var log = logs.NewLogger()

func init() {
	log.SetLogger(logs.AdapterConsole)
	log.Debug("this is a debug message")
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "marcelo@mail.com"
	c.TplName = "index.tpl"
}

type PersonFromForm struct {
	Name      string `form:"name"`
	Email     string `form:"email"`
	Phone_num string `form:"phone_number"`
}

func (c *MainController) Post() {
	// If the ajax is sent in the following way
	// Javascript snippet below
	// 	$.ajax({
	//            type: "POST",
	//            url: "/",
	//            data: {
	//		"name": $("#name").val(),
	//		"email": $("#email").val(),
	//		"phone_number": $("#phone_number").val(),
	//	    } // The data passed to / by POST method
	//        }).done(function(msg) {
	//	    console.log(msg)

	// then we can capture the values from each field like this
	//name      := c.GetString("name") // Get the data "name" from request
	//email     := c.GetString("email") // Get the data "email" from request
	//phone_num := c.GetString("phone_number") // Get the data "phone_number" from request
	//log.Info(name)
	//log.Info(email)
	//log.Info(phone_num)

	// Storing data in an anonymour struct
	// to be sent back to the browser (HTTP response) in JSON format
	//	result := struct {
	//		Name       string
	//		Email      string
	//		Phone_num  string
	//	}{name, email, phone_num}
	//c.Data["json"] = &result

	// Go OOP here
	// Faster way to get each form field and
	// assemble an instance of a Person object
	// with its respective attributes populated
	//	p := PersonFromForm{}
	//      //p := models.Person{} // Using class from other package
	// However this only works if the attributes are mapped with 'form:"name"' tags
	//	if err := c.ParseForm(&p); err != nil {
	//		log.Error("Error while processing request parameters")
	//	}

	p := models.Person{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &p)

	log.Info(p.Name)
	log.Info(p.Email)
	log.Info(p.Phone_num)

	c.Data["json"] = &p
	c.ServeJSON()
}
