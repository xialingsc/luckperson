package main

import (
	//	"fmt"
	//	"strings"
	//"fmt"
	"github.com/astaxie/beego"
	//	"luckperson/controllers"
	_ "luckperson/routers"
	//	"sort"
)

func main() {
	//var path string = "/Users/xialing/luckperson/src/luckperson/conf/test.txt"
	//controllers.ReadFile(path)
	//	fmt.Println(strings.EqualFold("Go", "go"))
	//	s := []string{"PHP", "golang", "python", "C", "Objective-C"}
	//	sort.Sort(sort.StringSlice(s))
	//	fmt.Println(s)
	beego.Run()
}
