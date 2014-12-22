package controllers

import (
	"MYProject/QandA/models"
	"encoding/xml"
	//"fmt"
	"github.com/astaxie/beego"
	"os"
)

type UPController struct {
	beego.Controller
}

type User struct {
	Name    string `form:"file_name,text,文件名:"`
	Imgfile string `form:"image_file,file,上传文件:"`
	Submit  string `form:",submit"`
}

func (this *UPController) Post() {
	list := this.GetString("list")
	o := models.List1
	switch list {
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
	default:
		return
	}

	qimgName := this.GetString("qimgName")
	if qimgName != "" {
		err := this.SaveToFile("qimg", "static/img/"+qimgName)
		if err != nil {
			this.Ctx.Output.Body([]byte(err.Error()))
		} else {
			this.Ctx.Output.Body([]byte("文件上传 OK"))
		}
	}
	qcontent := this.GetString("qcontent")
	content := []models.Content{models.Content{qcontent, qimgName}}
	key, err := this.GetInt("key")
	if err != nil {
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	qhead := this.GetString("qhead")
	n := models.Question{key, qhead, content}
	//fmt.Printf("%#v\n", n)
	o.Add(n)
	//fmt.Printf("%#v\n", o)
	f, _ := os.Create("list" + list + ".xml")
	defer f.Close()
	body, _ := xml.Marshal(o)
	f.Write(body)
	this.Ctx.Output.Body([]byte("生成XML OK"))
}
func (this *UPController) Get() {
	//fmt.Println(models.GetImg())
	//this.Data["json"] = models.GetImg()
	//this.ServeJson()
	//this.Data["Form"] = &User{Submit: "提交"}
	this.TplNames = "updown_1.tpl"
	//this.Ctx.Output.Body([]byte("please try post"))
}
