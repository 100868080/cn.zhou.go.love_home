package main

import (
	_ "2018/cn.zhou.beego/love_home/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//11100abc
	fmt.Println("abc,hello,world")
	beego.Run()
}
