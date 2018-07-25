package src

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello world")
}
func main() {
	beego.Router("/go_file/src", &HomeController{})
	beego.Run()
}
