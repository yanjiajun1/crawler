package main

import (
	"./engine"
    "./scheduler"
	"./zhenai/parser"
)
func main() {
    e:=engine.ConcurrentEngine{
    	Scheduler: &scheduler.SimpleScheduler{},
    	WorkerCount:10,
	}
    e.Run(engine.Request{Url:"http://www.zhenai.com/zhenghun",
    ParserFunc:parser.ParseCityList,
    })
}
