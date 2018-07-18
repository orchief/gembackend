package main

import (
	_ "github.com/astaxie/beego/config/xml"
	"github.com/astaxie/beego/config"
	"fmt"
	"github.com/gembackend/scripts"
)

func main() {
	iniconf, err := config.NewConfig("ini", "/Users/yinhaozheng/go/src/github.com/gembackend/startscript/test.yaml")
	fmt.Println(iniconf.String("demo::key2"))
	fmt.Println(err)
	//u := scripts.NewEthUpdaterWeb3(5000000)
	//g(u)
}

func g(u scripts.Updater) {
	u.Forever()
}
