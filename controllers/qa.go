package controllers

import (
	"MYProject/QandA/models"
	"fmt"
	"github.com/astaxie/beego"
)

type QAController struct {
	beego.Controller
}

func (this *QAController) Get() {
	id := this.Ctx.Input.Param(":id")
	fmt.Println(id)
	o := models.List1
	switch id {
	case "1":
		o = models.List1
	case "2":
		o = models.List2
	case "3":
		o = models.List3
	case "4":
		o = models.List4
	case "5":
		o = models.List5
	case "6":
		o = models.List6

	}
	//a := models.Content{"HHHH", "listimg1.png"}
	//b := models.Content{"HHHH", ""}
	//s := []models.Content{}
	//s = append(append(s, a), b)

	//n := models.Question{1, "quetion", s}
	////m := []models.Question{n}

	//o.Add(n)
	//fmt.Printf("\n%#v", models.List1)
	this.Data["Page"] = o
	this.TplNames = "list.tpl"

}
