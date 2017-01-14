package serverUtuls

import (
	"fmt"
	"encoding/json"
	"strings"
)

type Request struct {
	Site []string // Slice of strings: https://blog.golang.org/go-slices-usage-and-internals
	SearchText string
}
func (r *Request )init(_Site []string, _SearchText string){
	r.Site = _Site
	r.SearchText = _SearchText
}

type Response struct {
	FoundAtSite string
}
func checkError(msg string, e *error ){
	if *e != nil{
		fmt.Println(msg,*e)
	}
}
func decodeJSON(from [] byte, to interface{})error{
	return json.NewDecoder(strings.NewReader(string(from))).Decode(&to)
}
