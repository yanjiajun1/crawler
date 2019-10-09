package main

import (
	"./fetcher"
	"fmt"
	"regexp"
)
func main() {
	all,_:=fetcher.Fetch("http://www.zhenai.com/zhenghun")
    re:=regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*[^<]+</a>`)
    matches:=re.FindAll(all,-1)
    for _,m:=range matches{
    	fmt.Printf("%s\n",m)
	}
}
