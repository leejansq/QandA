package models

import (
	"encoding/xml"
	//"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type Content struct {
	Page string
	Img  string
}

type Question struct {
	Key   int64
	Qustn string
	Contn []Content
}

type Page struct {
	Head string
	Body []Question
}

type Qusort []Question

func (s Qusort) Len() int {
	return len(s)
}

func (s Qusort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Qusort) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

func (p *Page) Add(Q Question) {
	flag := false
	for k, v := range p.Body {
		if v.Key == Q.Key {
			p.Body[k] = Q
			flag = true
		}
	}
	if flag == false {
		p.Body = append(p.Body, Q)
		sort.Sort(Qusort(p.Body))
	}
	//fmt.Printf("%#v", p)

}

var (
	List1 = &Page{Head: "视频查看", Body: []Question{}}
	List2 = &Page{Head: "设备升级", Body: []Question{}}
	List3 = &Page{Head: "设备重置", Body: []Question{}}
	List4 = &Page{Head: "WIFI强度", Body: []Question{}}
	List5 = &Page{Head: "设备连接", Body: []Question{}}
	List6 = &Page{Head: "状态灯含义", Body: []Question{}}
)

func GetImg() (imgs []string) {
	filepath.Walk("./static/img",
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			imgs = append(imgs, f.Name())
			// println(path)
			return nil
		})
	return

}

func Ixml(s string, p *Page) {
	f, err := os.Open(s)
	if err == nil {
		defer f.Close()
		b, err1 := ioutil.ReadAll(f)
		if err1 == nil {
			xml.Unmarshal(b, p)
		}
	}
}

func init() {
	Ixml("list1.xml", List1)
	Ixml("list2.xml", List2)
	Ixml("list3.xml", List3)
	Ixml("list4.xml", List4)
	Ixml("list5.xml", List5)
	Ixml("list6.xml", List6)
}
