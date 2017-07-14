package main

import (
	_ "test/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	fmt.Println("test 1")
	beego.Run()
}

