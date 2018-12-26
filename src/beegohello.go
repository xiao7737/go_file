package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("Hello World!")
}
func main() {
	beego.Router("/go_file/src", &HomeController{})
	beego.Run()
}
