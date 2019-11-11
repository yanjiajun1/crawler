package parser

import (
	"../../engine"
	_ "fmt"
	"regexp"
)
const cityRe=`<a href="(http://album.zhenai.com/u/[1-9]*)" [^>]*>([^<]+)</a>`
func ParserCity(contents []byte) engine.ParserResult{
	re:=regexp.MustCompile(cityRe)
	matches:=re.FindAllSubmatch(contents,-1)
	result:=engine.ParserResult{}
	for _,m:=range matches{
		result.Items=append(result.Items ,"User"+string(m[2]))
//		log.Printf("user url %s",string(m[1]))
		result.Requests=append(result.Requests,engine.Request{Url:string(m[1]),ParserFunc:engine.NilParser})
		//fmt.Println("city %s , URL:%s\n",m[2],m[1])

	}
	return result
}