package engine
import(
	"../fetcher"
	"log"
)
type SimpleEngine struct {

}
func (SimpleEngine)Run(Seeds ...Request){
	var requests []Request
	for _,r:=range Seeds{
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
	//	log.Printf("Fetching %s",r.Url)
		requests=requests[1:]
		ParserResult,err:=Worker(r)
		if err !=nil{
			continue
		}
		requests=append(requests,ParserResult.Requests...)
		for _, item:=range ParserResult.Items{
			log.Printf("Got item %s",item)
		}
	}
}
func Worker(r Request) (ParserResult ,error) {
	body,err:=fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetcher:error"+"fetching url %s:%v",r.Url,err)
       return ParserResult{},err
	}
	return r.ParserFunc(body),nil
}
