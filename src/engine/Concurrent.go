package engine

import (
	"log"
)

type ConcurrentEngine struct {
   Scheduler  Scheduler
   WorkerCount int
}
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}
func (e *ConcurrentEngine) Run (seeds ...Request){

	in :=make(chan Request)
	out :=make(chan ParserResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i:=0 ;i<e.WorkerCount;i++{
		createWorker(in,out)
	}
	for _,r:= range seeds{
		e.Scheduler.Submit(r)
	}
	for {
		result :=<-out
		for _,item:=range result.Items{
			log.Printf("Go item %v",item)
		}
		for _,request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}

}
func createWorker(in chan Request,out chan ParserResult){
	go func() {
		for {
			request := <-in
			Worker(request)
			result, err := Worker(request)
			if err !=nil {
				continue
			}
			out <- result
		}
	}()
}