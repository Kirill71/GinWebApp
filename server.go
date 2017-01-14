package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
	"strings"
	"io/ioutil"
"csutils"
)

const BAD_RESPONSE = "HTTP Code 204 No Content"

type Request struct {
	Site []string
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
	serverUtuls.
}

func decodeJSON(from [] byte, to interface{})error{
	return json.NewDecoder(strings.NewReader(string(from))).Decode(&to)
}

func server(){
	gin.SetMode(gin.ReleaseMode)
	router:= gin.Default()
	router.POST("",checkTestHandler)
	router.Run(":8181")
}

func checkTestHandler(c *gin.Context){
	request := Request{}
	response := Response{}
	requestBody,_:= ioutil.ReadAll(c.Request.Body)
	err := decodeJSON(requestBody,request)
	checkError("json was not decode",&err)
	response.FoundAtSite = BAD_RESPONSE
	for _, str := range request.Site {
		if strings.Contains(strings.ToLower(str), strings.ToLower(request.SearchText)) {
			response.FoundAtSite = str
			break
		}
	}
	c.JSON(200, gin.H{
			"FoundAtSite": response.FoundAtSite,
		})
}


func main(){
	 server()
}
